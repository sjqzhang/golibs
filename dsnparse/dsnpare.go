package dsnparse

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type DSN struct {
	url *url.URL
	*DSNValues
}

// parses dsn string and returns DSN instance
func Parse(dsn string) (*DSN, error) {
	reg := regexp.MustCompile(`tcp\(.*?\)`) //uniform url format
	if m := reg.FindStringSubmatch(dsn); len(m) > 0 {
		match := m[0]
		match = strings.TrimPrefix(match, "tcp(")
		match = strings.TrimSuffix(match, ")")
		dsn = reg.ReplaceAllString(dsn, match)
	}
	parsed, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}
	d := DSN{
		parsed,
		&DSNValues{parsed.Query()},
	}
	return &d, nil
}

// Parses query and returns dsn values
func ParseQuery(query string) (*DSNValues, error) {
	parsed, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}
	return &DSNValues{parsed}, nil
}

// returns DSNValues from url.Values
func NewValues(query url.Values) (*DSNValues, error) {
	return &DSNValues{query}, nil
}

// return Host
func (d *DSN) HostWithPort() string {
	return d.url.Host
}

// return Host
func (d *DSN) Host() string {
	return strings.Split(d.url.Host, ":")[0]
}

// return Host
func (d *DSN) Port() string {
	hp := strings.Split(d.url.Host, ":")
	if len(hp) == 2 {
		return hp[1]
	} else {
		return ""
	}
}

// return Scheme
func (d *DSN) Scheme() string {
	return d.url.Scheme
}

// returns path
func (d *DSN) Path() string {
	return d.url.Path
}

// returns path
func (d *DSN) DatabaseName() string {
	return strings.Replace(d.url.Path, "/", "", -1)
}

// returns user
func (d *DSN) User() *url.Userinfo {
	return d.url.User
}

// returns Username
func (d *DSN) Username() string {
	return d.url.User.Username()
}

// returns Username
func (d *DSN) Password() string {
	v, ok := d.url.User.Password()
	if ok {
		return v
	} else {
		return ""
	}
}

// DSN Values
type DSNValues struct {
	url.Values
}

// returns int value
func (d *DSNValues) GetInt(paramName string, defaultValue int) int {
	value := d.Get(paramName)
	if i, err := strconv.Atoi(value); err == nil {
		return i
	} else {
		return defaultValue
	}
}

// returns string value
func (d *DSNValues) GetString(paramName string, defaultValue string) string {
	value := d.Get(paramName)
	if value == "" {
		return defaultValue
	} else {
		return value
	}
}

// returns string value
func (d *DSNValues) GetBool(paramName string, defaultValue bool) bool {
	value := strings.ToLower(d.Get(paramName))
	if value == "true" || value == "1" {
		return true
	} else if value == "0" || value == "false" {
		return false
	} else {
		return defaultValue
	}
}

// returns string value
func (d *DSNValues) GetSeconds(paramName string, defaultValue time.Duration) time.Duration {
	if i, err := strconv.Atoi(d.Get(paramName)); err == nil {
		return time.Duration(i) * time.Second
	} else {
		return defaultValue
	}
}
