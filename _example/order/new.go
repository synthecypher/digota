//     Digota <http://digota.com> - eCommerce microservice
//     Copyright (C) 2017  Yaron Sumel <yaron@digota.com>. All Rights Reserved.
//
//     This program is free software: you can redistribute it and/or modify
//     it under the terms of the GNU Affero General Public License as published
//     by the Free Software Foundation, either version 3 of the License, or
//     (at your option) any later version.
//
//     This program is distributed in the hope that it will be useful,
//     but WITHOUT ANY WARRANTY; without even the implied warranty of
//     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//     GNU Affero General Public License for more details.
//
//     You should have received a copy of the GNU Affero General Public License
//     along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"github.com/synthecypher/digota/order/orderpb"
	"github.com/synthecypher/digota/payment/paymentpb"
	"github.com/synthecypher/digota/sdk"
	"golang.org/x/net/context"
	"log"
)

func main() {

	c, err := sdk.NewClient("localhost:3051", &sdk.ClientOpt{
		InsecureSkipVerify: false,
		ServerName:         "server.com",
		CaCrt:              "out/ca.crt",
		Crt:                "out/client.com.crt",
		Key:                "out/client.com.key",
	})

	if err != nil {
		panic(err)
	}

	defer c.Close()

	// Create new order
	o, err := orderpb.NewOrderServiceClient(c).New(context.Background(), &orderpb.NewRequest{
		Currency: paymentpb.Currency_USD,
		Items: []*orderpb.OrderItem{
			{
				Parent:   "af350ecc-56c8-485f-8858-74d4faffa9cb",
				Quantity: 2,
				Type:     orderpb.OrderItem_sku,
			},
			{
				Parent:   "af350ecc-56c8-485f-8858-74d4faffa9cb",
				Quantity: 2,
				Type:     orderpb.OrderItem_sku,
			},
			//{
			//	Parent:   "480e53bf-b409-4a34-8c74-13786b35ae11",
			//	Quantity: 1,
			//	Type:     orderpb.OrderItem_sku,
			//},
			//{
			//	Parent:   "480e53bf-b409-4a34-8c74-13786b35ae11",
			//	Quantity: 1,
			//	Type:     orderpb.OrderItem_sku,
			//},
			{
				Amount:      -1000,
				Description: "on the fly discount without parent",
				Currency:    paymentpb.Currency_USD,
				Type:        orderpb.OrderItem_discount,
			},
			{
				Amount:      1000,
				Description: "Tax (Included)",
				Currency:    paymentpb.Currency_USD,
				Type:        orderpb.OrderItem_tax,
			},
		},
		Email: "yaron@digota.com",
		Shipping: &orderpb.Shipping{
			Name:  "Yaron Sumel",
			Phone: "+972 000 000 000",
			Address: &orderpb.Shipping_Address{
				Line1:      "Loren ipsum",
				City:       "San Jose",
				Country:    "USA",
				Line2:      "",
				PostalCode: "12345",
				State:      "CA",
			},
		},
	})

	if err != nil {
		panic(err)
	}

	log.Println(o.GetId())

}
