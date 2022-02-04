package restapi

import (
	"fmt"
	"net/url"
)

type Restapi struct {
	Basepath url.URL
	Client   *HttpClient
}

func NewRestapi(base string) (*Restapi, error) {
	rp := Restapi{}
	u, er := url.Parse(base)
	if er != nil {
		return &rp, er
	}
	rp.Basepath = *u

	c := NewHttpClient()

	rp.Client = &c

	return &rp, nil
}

//Parse Full Urls using base paths

//Get Page URL
func (rp *Restapi) getpageurl(pagname string) (string, error) {
	pageaddr := "/page/" + pagname
	u, er := rp.Basepath.Parse(pageaddr)
	if er != nil {
		return "", er
	}

	return u.String(), nil
}

func (rp *Restapi) getorgurl(path string) (string, error) {
	pageaddr := "/org/" + path
	u, er := rp.Basepath.Parse(pageaddr)
	if er != nil {
		return "", er
	}

	return u.String(), nil
}

func (rp *Restapi) getweburl(path string) (string, error) {
	pageaddr := "/web/" + path
	u, er := rp.Basepath.Parse(pageaddr)
	if er != nil {
		return "", er
	}

	return u.String(), nil
}

// Actual Functions to get and post data

func (rp *Restapi) GetPage(pagename string) (string, error) {
	pageaddr, err := rp.getpageurl(pagename)
	if err != nil {
		return "", err
	}

	dat, err := rp.Client.Get(pageaddr)

	return string(dat), err
}

func (rp *Restapi) PostPage(pagename string, dat []byte, append bool) (bool, error) {
	pageaddr, err := rp.getpageurl(pagename)
	if err != nil {
		return false, err
	}

	if append {
		stat, err := rp.Client.PostAppend(pageaddr, dat)
		return stat, err
	} else {
		stat, err := rp.Client.Post(pageaddr, dat)
		return stat, err
	}

}

func (rp *Restapi) Workingsub() (string, error) {
	pageaddr, err := rp.getweburl("/working")
	if err != nil {
		return "", err
	}

	dat, er := rp.Client.Get(pageaddr)
	if er != nil {
		return "", er
	}

	return string(dat), nil
}

//Get Checklist pretty version
//get checklist web or org
func (rp *Restapi) GetCheckList(web bool) ([]byte, error) {
	var pageaddr string
	var err error

	if web {
		pageaddr, err = rp.getweburl("/checklist")
		if err != nil {
			return []byte{}, err
		}
	} else {
		pageaddr, err = rp.getorgurl("/checklist")
		if err != nil {
			return []byte{}, err
		}
	}

	dat, er := rp.Client.Get(pageaddr)
	if er != nil {
		return []byte{}, er
	}

	return dat, nil

}

//save to db
func (rp *Restapi) SavetoDB() error {
	pageaddr := "/commit"
	u, er := rp.Basepath.Parse(pageaddr)
	if er != nil {
		return er
	}

	_, err := rp.Client.Get(u.String())
	if err != nil {
		return err
	}
	return nil
}

//web tool output
func (rp *Restapi) GetWebToolOutput(sub string, toolname string) (string, error) {
	pageaddr, err := rp.getweburl(sub + "/" + toolname)
	fmt.Println(pageaddr)
	if err != nil {
		return "", err
	}

	dat, err := rp.Client.Get(pageaddr)

	return string(dat), err
}

//org tool output
func (rp *Restapi) GetOrgToolOutput(toolname string) (string, error) {
	pageaddr, err := rp.getorgurl(toolname)
	if err != nil {
		return "", err
	}

	dat, err := rp.Client.Get(pageaddr)

	return string(dat), err
}

//post web tool output
func (rp *Restapi) PostWebToolOutput(sub string, toolname string, dat []byte, append bool) (bool, error) {
	pageaddr, err := rp.getweburl(sub + "/" + toolname)
	if err != nil {
		return false, err
	}

	if append {
		stat, err := rp.Client.PostAppend(pageaddr, dat)
		return stat, err
	} else {
		stat, err := rp.Client.Post(pageaddr, dat)
		return stat, err
	}

}

//post org tool output
func (rp *Restapi) PostOrgToolOutput(toolname string, dat []byte, append bool) (bool, error) {
	pageaddr, err := rp.getorgurl(toolname)
	if err != nil {
		return false, err
	}

	if append {
		stat, err := rp.Client.PostAppend(pageaddr, dat)
		return stat, err
	} else {
		stat, err := rp.Client.Post(pageaddr, dat)
		return stat, err
	}

}

func (rp *Restapi) Alive() {
	scode, err := rp.Client.StatusCode(rp.Basepath.String())
	if err != nil {
		fmt.Println("It looks like BBSidekick App is Not Running")
		fmt.Println("Or is Using A Different Port Use url if that's the case")
		panic(err)
	}
	if scode != 200 {
		panic(fmt.Errorf("server is down or returned %v status code", scode))
	}

}
