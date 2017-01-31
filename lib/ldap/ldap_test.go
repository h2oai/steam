package ldap

import (
	"crypto/tls"
	"os"
	"reflect"
	"testing"

	"github.com/go-ldap/ldap"
)

var (
	adPass string
)

func init() {
	adPass = os.Getenv("AD_PASSWORD")
}

func TestLdap_TestConfig(t *testing.T) {
	type fields struct {
		Address                 string
		BindDN                  string
		BindPass                string
		UserBaseDn              string
		UserBaseFilter          string
		UserNameAttribute       string
		GroupBaseDn             string
		GroupNameAttribute      string
		GroupNames              []string
		StaticGroupSearchFilter string
		StaticMemberAttribute   string
		SearchRequestSizeLimit  int
		SearchRequestTimeLimit  int
		Ldaps                   bool
		ForceBind               bool
		tlsConfig               *tls.Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		want1   map[string]int
		wantErr bool
	}{
		{
			name: "basic test",
			fields: fields{
				Address:               "ldap.0xdata.loc:389",
				BindDN:                "cn=admin,dc=0xdata,dc=loc",
				BindPass:              "0xdata",
				UserBaseDn:            "ou=users,dc=0xdata,dc=loc",
				UserNameAttribute:     "uid",
				GroupBaseDn:           "ou=groups,dc=0xdata,dc=loc",
				GroupNameAttribute:    "cn",
				StaticMemberAttribute: "memberUid",
				GroupNames:            []string{"jettygroup", "steamgroup", "randomgroup"},
				Ldaps:                 false,
			},
			want: 8,
			want1: map[string]int{
				"jettygroup": 8,
			},
			wantErr: false,
		}, {
			name: "AD_test",
			fields: fields{
				Address:               "ad.h2o.ai:389",
				BindDN:                "CN=Test,CN=Users,DC=ad,DC=h2o,DC=ai",
				BindPass:              adPass,
				UserBaseDn:            "CN=Users,DC=ad,DC=h2o,DC=ai",
				UserNameAttribute:     "sAMAccountName",
				GroupBaseDn:           "OU=engineering,DC=ad,DC=h2o,DC=ai",
				GroupNameAttribute:    "CN",
				StaticMemberAttribute: "member",
				GroupNames:            []string{"data science", "steamgroup"},
			},
			want: 2,
			want1: map[string]int{
				"data science": 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Ldap{
				Address:                 tt.fields.Address,
				BindDN:                  tt.fields.BindDN,
				BindPass:                tt.fields.BindPass,
				UserBaseDn:              tt.fields.UserBaseDn,
				UserBaseFilter:          tt.fields.UserBaseFilter,
				UserNameAttribute:       tt.fields.UserNameAttribute,
				GroupBaseDn:             tt.fields.GroupBaseDn,
				GroupNameAttribute:      tt.fields.GroupNameAttribute,
				GroupNames:              tt.fields.GroupNames,
				StaticGroupSearchFilter: tt.fields.StaticGroupSearchFilter,
				StaticMemberAttribute:   tt.fields.StaticMemberAttribute,
				SearchRequestSizeLimit:  tt.fields.SearchRequestSizeLimit,
				SearchRequestTimeLimit:  tt.fields.SearchRequestTimeLimit,
				Ldaps:     tt.fields.Ldaps,
				ForceBind: tt.fields.ForceBind,
				tlsConfig: tt.fields.tlsConfig,
			}
			got, got1, err := l.TestConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("Ldap.TestConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Ldap.TestConfig() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Ldap.TestConfig() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLdap_checkGroup(t *testing.T) {
	adConn, err := ldap.Dial("tcp", "ad.h2o.ai:389")
	if err != nil {
		t.Skipf("no AD connection found")
	}
	defer adConn.Close()
	err = adConn.Bind("CN=Test,CN=Users,DC=ad,DC=h2o,DC=ai", adPass)
	if err != nil {
		t.Skipf("invalid bind")
	}

	conn, err := ldap.Dial("tcp", "ldap.0xdata.loc:389")
	if err != nil {
		t.Skipf("no ldap connection found")
	}
	defer conn.Close()

	type fields struct {
		Address                 string
		BindDN                  string
		BindPass                string
		UserBaseDn              string
		UserBaseFilter          string
		UserNameAttribute       string
		GroupBaseDn             string
		GroupNameAttribute      string
		GroupNames              []string
		StaticGroupSearchFilter string
		StaticMemberAttribute   string
		SearchRequestSizeLimit  int
		SearchRequestTimeLimit  int
		Ldaps                   bool
		ForceBind               bool
		tlsConfig               *tls.Config
	}
	type args struct {
		conn   *ldap.Conn
		userCN string
		userDN string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "basic_test",
			fields: fields{
				UserBaseDn:            "ou=users,dc=0xdata,dc=loc",
				UserNameAttribute:     "uid",
				GroupBaseDn:           "ou=groups,dc=0xdata,dc=loc",
				GroupNameAttribute:    "cn",
				StaticMemberAttribute: "memberUid",
				GroupNames:            []string{"jettygroup", "steamgroup", "randomgroup"},
			},
			args: args{
				conn:   conn,
				userCN: "seb",
				userDN: "cn=seb,ou=users,dc=0xdata,dc=loc",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "AD_test",
			fields: fields{
				UserBaseDn:            "CN=Users,DC=ad,DC=h2o,DC=ai",
				UserNameAttribute:     "sAMAccountName",
				GroupBaseDn:           "OU=engineering,DC=ad,DC=h2o,DC=ai",
				GroupNameAttribute:    "CN",
				StaticMemberAttribute: "member",
				GroupNames:            []string{"data science", "steamgroup"},
			},
			args: args{
				conn:   adConn,
				userCN: "datascientist",
				userDN: "CN=datascientist,CN=Users,DC=ad,DC=h2o,DC=ai",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Ldap{
				Address:                 tt.fields.Address,
				BindDN:                  tt.fields.BindDN,
				BindPass:                tt.fields.BindPass,
				UserBaseDn:              tt.fields.UserBaseDn,
				UserBaseFilter:          tt.fields.UserBaseFilter,
				UserNameAttribute:       tt.fields.UserNameAttribute,
				GroupBaseDn:             tt.fields.GroupBaseDn,
				GroupNameAttribute:      tt.fields.GroupNameAttribute,
				GroupNames:              tt.fields.GroupNames,
				StaticGroupSearchFilter: tt.fields.StaticGroupSearchFilter,
				StaticMemberAttribute:   tt.fields.StaticMemberAttribute,
				SearchRequestSizeLimit:  tt.fields.SearchRequestSizeLimit,
				SearchRequestTimeLimit:  tt.fields.SearchRequestTimeLimit,
				Ldaps:     tt.fields.Ldaps,
				ForceBind: tt.fields.ForceBind,
				tlsConfig: tt.fields.tlsConfig,
			}
			got, err := l.checkGroup(tt.args.conn, tt.args.userCN, tt.args.userDN)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ldap.checkGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Ldap.checkGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLdap_CheckBind(t *testing.T) {
	type fields struct {
		Address                 string
		BindDN                  string
		BindPass                string
		UserBaseDn              string
		UserBaseFilter          string
		UserNameAttribute       string
		GroupBaseDn             string
		GroupNameAttribute      string
		GroupNames              []string
		StaticGroupSearchFilter string
		StaticMemberAttribute   string
		SearchRequestSizeLimit  int
		SearchRequestTimeLimit  int
		Ldaps                   bool
		ForceBind               bool
		tlsConfig               *tls.Config
	}
	type args struct {
		user     string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "AD_test",
			fields: fields{
				Address:               "ad.h2o.ai:389",
				BindDN:                "CN=Test,CN=Users,DC=ad,DC=h2o,DC=ai",
				BindPass:              adPass,
				UserBaseDn:            "CN=Users,DC=ad,DC=h2o,DC=ai",
				UserNameAttribute:     "sAMAccountName",
				GroupBaseDn:           "OU=engineering,DC=ad,DC=h2o,DC=ai",
				GroupNameAttribute:    "CN",
				StaticMemberAttribute: "member",
				GroupNames:            []string{"data science", "steamgroup"},
			},
			args: args{
				user: "datascientist",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Ldap{
				Address:                 tt.fields.Address,
				BindDN:                  tt.fields.BindDN,
				BindPass:                tt.fields.BindPass,
				UserBaseDn:              tt.fields.UserBaseDn,
				UserBaseFilter:          tt.fields.UserBaseFilter,
				UserNameAttribute:       tt.fields.UserNameAttribute,
				GroupBaseDn:             tt.fields.GroupBaseDn,
				GroupNameAttribute:      tt.fields.GroupNameAttribute,
				GroupNames:              tt.fields.GroupNames,
				StaticGroupSearchFilter: tt.fields.StaticGroupSearchFilter,
				StaticMemberAttribute:   tt.fields.StaticMemberAttribute,
				SearchRequestSizeLimit:  tt.fields.SearchRequestSizeLimit,
				SearchRequestTimeLimit:  tt.fields.SearchRequestTimeLimit,
				Ldaps:     tt.fields.Ldaps,
				ForceBind: tt.fields.ForceBind,
				tlsConfig: tt.fields.tlsConfig,
			}
			if err := l.CheckBind(tt.args.user, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("Ldap.CheckBind() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFromDatabase(t *testing.T) {
	type args struct {
		config    string
		tlsConfig *tls.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *Ldap
		wantErr bool
	}{
		{
			name: "basic test",
			args: args{
				config:    `{\"Address\":\"ldap.0xdata.loc:389\",\"Bind\":\"Y249YWRtaW4sZGM9MHhkYXRhLGRjPWxvYzoweGRhdGE=\",\"UserBaseDn\":\"ou=users,dc=0xdata,dc=loc\",\"UserBaseFilter\":\"\",\"UserNameAttribute\":\"uid\",\"GroupBaseDN\":\"ou=groups,dc=0xdata,dc=loc\",\"GroupNameAttribute\":\"cn\",\"GroupNames\":[\"jettygroup\",\"steamgroup\"],\"StaticMemberAttribute\":\"memberUid\",\"SearchRequestSizeLimit\":0,\"SearchRequestTimeLimit\":0,\"ForceBind\":false,\"Ldaps\":false}`,
				tlsConfig: &tls.Config{InsecureSkipVerify: true},
			},
			want: &Ldap{
				Address:               "ldap.0xdata.loc:389",
				BindDN:                "cn=admin,dc=0xdata,dc=loc",
				BindPass:              "0xdata",
				UserBaseDn:            "ou=users,dc=0xdata,dc=loc",
				UserNameAttribute:     "uid",
				GroupBaseDn:           "ou=groups,dc=0xdata,dc=loc",
				GroupNameAttribute:    "cn",
				StaticMemberAttribute: "memberUid",
				GroupNames:            []string{"jettygroup", "steamgroup"},
				Ldaps:                 false,
				tlsConfig:             &tls.Config{InsecureSkipVerify: true},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromDatabase(tt.args.config, tt.args.tlsConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}
