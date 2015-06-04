package bosh

import (
	"fmt"
//	"log"
	"os"
	"regexp"
	"strings"
	
	"golang.org/x/net/context"
	"github.com/codeskyblue/go-sh"
)

type BoshClient struct {

	Target string
	User string
	Password string
	
	UUID string
	Name string
	Version string
	CPI string
	
	Verbose bool
}

const boshCLI = "bosh %s --non-interactive --target %s --user %s --password %s --deployment %s %s"

func NewBoshClient(ctx context.Context, target string, user string, password string, verbose bool) (*BoshClient, error) {
	
	c := &BoshClient {
		Target: target,
		User: user,
		Password: password,
		
		Verbose: verbose,
	}
	
	return c, c.Connect()
}

func (b *BoshClient) Connect() error {
	
	result, err := b.exec("status", "")
	if err != nil {
		return err
	}
	
	re := regexp.MustCompile(
		"\\s+Name\\s+(.*)\\s+.*" + 
		"\\s+Version\\s+([.0-9]+)\\s+.*" + 
		"\\s+User\\s+.*" + 
		"\\s+UUID\\s+([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})" +
		"\\s+CPI\\s+(\\w+).*" )
	
	f := re.FindStringSubmatch(result)
	
	if len(f) > 0 {
		b.UUID = f[3]
		b.Name = f[1]
		b.Version = f[2]
		b.CPI = f[4]
	} else {
		if strings.Contains(result, "cannot access director") {
			return fmt.Errorf("bosh director not found")
		} else {
			return fmt.Errorf("bosh status did not get a valid response from a director at '%s':\n%s", b.Target, result)
		}
	}	
	return nil
}

func (b *BoshClient) IsConnected() bool {
	return b.UUID != ""
}

func (b *BoshClient) exec(cmd string, deployment string) (string, error) {
	
	session := sh.NewSession()
	
	args := []interface{}{ "--target", b.Target, "--user", b.User, "--password", b.Password }
	if b.Verbose {
		session.ShowCMD = true
		args = append(args, "--verbose")
	}
	if deployment != "" {
		if _, err := os.Stat(deployment); os.IsNotExist(err) {
	    	return "", fmt.Errorf("deployment file '%s' does not exist", deployment)
	    }
		args = append(args, "--deployment", deployment)		
	}
	for _, a := range strings.Fields(cmd){
		args = append(args, a)
	}
	
	out, err := session.Command("bosh", args...).Output()
	if err != nil {
		return "", err
	}	
	
	return string(out), nil
}
