package drivers

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
)

type SystemService struct {
	Name string
	Desc string
	Cmd string
}

func (s SystemService) Install()  {
	b := bytes.Buffer{}
	b.WriteString("[Unit]\n")
	b.WriteString(fmt.Sprintf("Description=%s\n",s.Desc))
	b.WriteString("After=network-online.target\n")
	b.WriteString("Wants=network-online.target\n")
	b.WriteString("[Service]\n")
	b.WriteString("Type=notify\n")
	b.WriteString(fmt.Sprintf("ExecStart=%s\n",s.Cmd))
	b.WriteString("ExecReload=/bin/kill -s HUP $MAINPID\n")
	b.WriteString("TimeoutSec=0\n")
	b.WriteString("RestartSec=2\n")
	b.WriteString("Restart=always\n")
	b.WriteString("StartLimitBurst=3\n")
	b.WriteString("StartLimitInterval=60s\n")
	b.WriteString("LimitNOFILE=infinity\n")
	b.WriteString("LimitNPROC=infinity\n")
	b.WriteString("LimitCORE=infinity\n")
	b.WriteString("TasksMax=infinity\n")
	b.WriteString("Delegate=yes\n")
	b.WriteString("KillMode=process\n")
	b.WriteString("OOMScoreAdjust=-500\n")
	b.WriteString("[Install]\n")
	b.WriteString("WantedBy=multi-user.target\n")
	ioutil.WriteFile(fmt.Sprintf("/usr/lib/systemd/system/%s.service",s.Name),b.Bytes(),fs.ModePerm)
	//daemon-reload
	c := &exec.Cmd{
		Path: "/bin/systemctl",
		Args: []string{"daemon-reload"},
	}
	c.Run()
	s.Enable()
	s.Start()
}
func (s SystemService) UnInstall()  {
	s.Stop()
	c := &exec.Cmd{
		Path: "/bin/systemctl",
		Args: []string{"disable",s.Name + ".service"},
	}
	c.Run()
	os.Remove(fmt.Sprintf("/usr/lib/systemd/system/%s.service",s.Name))
	c = &exec.Cmd{
		Path: "/bin/systemctl",
		Args: []string{"daemon-reload"},
	}
	c.Run()


}
func (s SystemService) Restart()  {
	c := &exec.Cmd{
		Path: "/bin/systemctl",
		Args: []string{"restart",s.Name + ".service"},
	}
	c.Run()
}
func (s SystemService) Enable()  {
	c := &exec.Cmd{
		Path: "/bin/systemctl",
		Args: []string{"enable",s.Name + ".service"},
	}
	c.Run()
}
func (s SystemService) Stop()  {
	c := &exec.Cmd{
		Path: "/bin/systemctl",
		Args: []string{"stop",s.Name + ".service"},
	}
	c.Run()
}
func (s SystemService) Start()  {
	c := &exec.Cmd{
		Path: "/bin/systemctl",
		Args: []string{"start",s.Name + ".service"},
	}
	c.Run()
}
