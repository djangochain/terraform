// provider.go
package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var ()

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"example_server": resourceServer(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(context context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	// Warning or errors can be collected in a slice type
	fmt.Println("username", username)
	fmt.Println("password", password)

	// if (username != "") && (password != "") {
	//   c, err := hashicups.NewClient(nil, &username, &password)
	//   if err != nil {
	// 	return nil, diag.FromErr(err)
	//   }

	//   return c, diags
	// }

	// c, err := hashicups.NewClient(nil, nil, nil)
	// if err != nil {
	//   return nil, diag.FromErr(err)
	// }
	token := username + "-" + password
	return token, nil
}
