package drivers

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os/user"
	"strings"
)

type CrontabDriver struct {
	IsLock bool
	Cmd string
	Timer string
}

func (c CrontabDriver) ListAll()  {
	u,_:=user.Current()
	b,e := ioutil.ReadFile(fmt.Sprintf("/var/spool/cron/%s",u.Username))
	if e != nil{
		fmt.Println("当前用户没有计划任务")
	}else{
		fmt.Println(string(b))
	}
}
func (c CrontabDriver) Current()  {
	u,_:=user.Current()
	b,e := ioutil.ReadFile(fmt.Sprintf("/var/spool/cron/%s",u.Username))
	if e != nil{
		fmt.Println("当前用户没有计划任务")
	}else{
		cs := string(b)
		lines := strings.Split(cs,"\n")
		for i,line := range lines{
			if strings.Index(line,c.Cmd) > 0{
				fmt.Sprintf("%d -> %s\n",i,line)
			}
		}
	}
}
func (c CrontabDriver) Save(){
	u,_:=user.Current()
	news := []string{}
	cf := fmt.Sprintf("/var/spool/cron/%s",u.Username)
	b,e := ioutil.ReadFile(cf)
	if e == nil{
		cs := string(b)
		lines := strings.Split(cs,"\n")
		for _,line := range lines{
			if strings.Index(line,c.Cmd) < 0{
				news = append(news,line)
			}
		}
	}
	if c.IsLock{
		news = append(news,fmt.Sprintf("%s flock -xn %s.lock -c '%s'",c.Timer,c.Cmd,c.Cmd))
	}else{
		news = append(news,fmt.Sprintf("%s %s",c.Timer,c.Cmd))
	}
	ioutil.WriteFile(cf,[]byte(strings.Join(news,"\n")),fs.ModePerm)
}
