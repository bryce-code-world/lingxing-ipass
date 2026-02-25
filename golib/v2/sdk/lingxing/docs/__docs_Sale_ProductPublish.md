刊登管理-提交商品资料
接口信息
API Path	请求协议	请求方式	令牌桶容量
/listing/publish/openapi/amazon/product/publish	HTTPS	POST	1
请求参数
参数名	说明	必填	类型	示例
store_id	store_id	是	[number]	3690
data		是	[array]	
data>>sku	sku	是	[string]	abced-1
data>>productType	商品类型	是	[string]	LOCK
data>>attributes	商品属性对象	是	[object]	
请求示例
{
    "store_id": 123,
    "data": [
        {
            "sku": "testopenapi-01-P",
            "productType": "LOCK",
            "attributes": {
                "item_name": [
                    {
                        "value": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "brand": [
                    {
                        "value": "Generic",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "supplier_declared_has_product_identifier_exemption": [
                    {
                        "value": true,
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "item_type_keyword": [
                    {
                        "value": "combination-padlocks",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "model_number": [
                    {
                        "value": "LX-test",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "model_name": [
                    {
                        "value": "LX-test",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "manufacturer": [
                    {
                        "value": "GENERIC",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "fulfillment_availability": [
                    {
                        "fulfillment_channel_code": "AMAZON_NA"
                    }
                ],
                "purchasable_offer": [
                    {
                        "currency": "USD",
                        "marketplace_id": "ATVPDKIKX0DER",
                        "our_price": [
                            {
                                "schedule": [
                                    {
                                        "value_with_tax": "9.99"
                                    }
                                ]
                            }
                        ],
                        "audience": "ALL"
                    }
                ],
                "condition_type": [
                    {
                        "value": "new_new",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "list_price": [
                    {
                        "value": "11.99",
                        "currency": "USD",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "product_description": [
                    {
                        "value": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "bullet_point": [
                    {
                        "value": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "generic_keyword": [
                    {
                        "value": "combination lock",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    },
                    {
                        "value": "lock",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "special_feature": [
                    {
                        "value": "Anti-Drill",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "style": [
                    {
                        "value": "modem",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "number_of_items": [
                    {
                        "value": "2",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "number_of_pieces": [
                    {
                        "value": "1",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "lock_type": [
                    {
                        "value": "Combination Lock",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "included_components": [
                    {
                        "value": "two lock body",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "country_of_origin": [
                    {
                        "value": "CN",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "batteries_required": [
                    {
                        "value": false,
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "supplier_declared_dg_hz_regulation": [
                    {
                        "value": "not_applicable",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "item_dimensions": [
                    {
                        "length": {
                            "value": "1.8",
                            "unit": "inches"
                        },
                        "width": {
                            "value": "0.8",
                            "unit": "inches"
                        },
                        "height": {
                            "value": "3.5",
                            "unit": "inches"
                        },
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "item_package_dimensions": [
                    {
                        "length": {
                            "value": "2",
                            "unit": "inches"
                        },
                        "width": {
                            "value": "1",
                            "unit": "inches"
                        },
                        "height": {
                            "value": "4",
                            "unit": "inches"
                        },
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "item_package_weight": [
                    {
                        "value": "0.5",
                        "unit": "pounds",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "parentage_level": [
                    {
                        "value": "parent",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "child_parent_sku_relationship": [
                    {
                        "child_relationship_type": "variation",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "variation_theme": [
                    {
                        "name": "COLOR"
                    }
                ]                
            }
        },
        {
            "sku": "testopenapi-01-C1",
            "productType": "LOCK",
            "attributes": {
                "item_name": [
                    {
                        "value": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "brand": [
                    {
                        "value": "Generic",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "supplier_declared_has_product_identifier_exemption": [
                    {
                        "value": true,
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "item_type_keyword": [
                    {
                        "value": "combination-padlocks",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "model_number": [
                    {
                        "value": "LX-test-C1",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "model_name": [
                    {
                        "value": "LX-test",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "manufacturer": [
                    {
                        "value": "GENERIC",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "fulfillment_availability": [
                    {
                        "fulfillment_channel_code": "AMAZON_NA"
                    }
                ],
                "purchasable_offer": [
                    {
                        "currency": "USD",
                        "marketplace_id": "ATVPDKIKX0DER",
                        "our_price": [
                            {
                                "schedule": [
                                    {
                                        "value_with_tax": "9.99"
                                    }
                                ]
                            }
                        ],
                        "audience": "ALL"
                    }
                ],
                "condition_type": [
                    {
                        "value": "new_new",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "list_price": [
                    {
                        "value": "11.99",
                        "currency": "USD",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "product_description": [
                    {
                        "value": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "bullet_point": [
                    {
                        "value": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "generic_keyword": [
                    {
                        "value": "combination lock",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    },
                    {
                        "value": "lock",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "special_feature": [
                    {
                        "value": "Anti-Drill",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "style": [
                    {
                        "value": "modem",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "number_of_items": [
                    {
                        "value": "2",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "color": [
                    {
                        "value": "SLIVER",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "number_of_pieces": [
                    {
                        "value": "1",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "lock_type": [
                    {
                        "value": "Combination Lock",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "included_components": [
                    {
                        "value": "two lock body",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "country_of_origin": [
                    {
                        "value": "CN",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "batteries_required": [
                    {
                        "value": false,
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "supplier_declared_dg_hz_regulation": [
                    {
                        "value": "not_applicable",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "item_dimensions": [
                    {
                        "length": {
                            "value": "1.8",
                            "unit": "inches"
                        },
                        "width": {
                            "value": "0.8",
                            "unit": "inches"
                        },
                        "height": {
                            "value": "3.5",
                            "unit": "inches"
                        },
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "item_package_dimensions": [
                    {
                        "length": {
                            "value": "2",
                            "unit": "inches"
                        },
                        "width": {
                            "value": "1",
                            "unit": "inches"
                        },
                        "height": {
                            "value": "4",
                            "unit": "inches"
                        },
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "item_package_weight": [
                    {
                        "value": "0.5",
                        "unit": "pounds",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "parentage_level": [
                    {
                        "value": "child",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "child_parent_sku_relationship": [
                    {
                        "child_relationship_type": "variation",
                        "parent_sku":"testopenapi-01-P",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "variation_theme": [
                    {
                        "name": "COLOR"
                    }
                ]                
            }
        },
        {
            "sku": "testopenapi-01-C2",
            "productType": "LOCK",
            "attributes": {
                "item_name": [
                    {
                        "value": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "brand": [
                    {
                        "value": "Generic",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "supplier_declared_has_product_identifier_exemption": [
                    {
                        "value": true,
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "item_type_keyword": [
                    {
                        "value": "combination-padlocks",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "model_number": [
                    {
                        "value": "LX-test-C2",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "model_name": [
                    {
                        "value": "LX-test",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "manufacturer": [
                    {
                        "value": "GENERIC",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "fulfillment_availability": [
                    {
                        "fulfillment_channel_code": "AMAZON_NA"
                    }
                ],
                "purchasable_offer": [
                    {
                        "currency": "USD",
                        "marketplace_id": "ATVPDKIKX0DER",
                        "our_price": [
                            {
                                "schedule": [
                                    {
                                        "value_with_tax": "9.99"
                                    }
                                ]
                            }
                        ],
                        "audience": "ALL"
                    }
                ],
                "condition_type": [
                    {
                        "value": "new_new",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "list_price": [
                    {
                        "value": "11.99",
                        "currency": "USD",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "product_description": [
                    {
                        "value": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "bullet_point": [
                    {
                        "value": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "generic_keyword": [
                    {
                        "value": "combination lock",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    },
                    {
                        "value": "lock",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "special_feature": [
                    {
                        "value": "Anti-Drill",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "style": [
                    {
                        "value": "modem",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "number_of_items": [
                    {
                        "value": "2",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "color": [
                    {
                        "value": "GREEN",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "number_of_pieces": [
                    {
                        "value": "1",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "lock_type": [
                    {
                        "value": "Combination Lock",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "included_components": [
                    {
                        "value": "two lock body",
                        "language_tag": "en_US",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "country_of_origin": [
                    {
                        "value": "CN",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "batteries_required": [
                    {
                        "value": false,
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "supplier_declared_dg_hz_regulation": [
                    {
                        "value": "not_applicable",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "item_dimensions": [
                    {
                        "length": {
                            "value": "1.8",
                            "unit": "inches"
                        },
                        "width": {
                            "value": "0.8",
                            "unit": "inches"
                        },
                        "height": {
                            "value": "3.5",
                            "unit": "inches"
                        },
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "item_package_dimensions": [
                    {
                        "length": {
                            "value": "2",
                            "unit": "inches"
                        },
                        "width": {
                            "value": "1",
                            "unit": "inches"
                        },
                        "height": {
                            "value": "4",
                            "unit": "inches"
                        },
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "item_package_weight": [
                    {
                        "value": "0.5",
                        "unit": "pounds",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "parentage_level": [
                    {
                        "value": "child",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "child_parent_sku_relationship": [
                    {
                        "child_relationship_type": "variation",
                        "parent_sku":"testopenapi-01-P",
                        "marketplace_id": "ATVPDKIKX0DER"
                    }
                ],
                "variation_theme": [
                    {
                        "name": "COLOR"
                    }
                ]                
            }
        }       
    ]
}
复制
错误
复制成功
返回结果
参数名	说明	必填	类型	示例
code	状态码	是	[number]	1
msg	消息提示	是	[string]	成功
data	响应数据	是	[object]	
data>>record_unique_id	批次唯一ID	是	[string]	613461911446487040
request_id	请求链路id	是	[string]	0fb7e663-3db9-4f6e-aef9-86c79aaf3ed6
返回成功示例
{
    "code": 1,
    "msg": "成功",
    "data": {
        "record_unique_id": "613461911446487040"
    },
    "request_id": "0fb7e663-3db9-4f6e-aef9-86c79aaf3ed6|013a0998-7dbd-11f0-86c8-4a2180c1ec00"
}
复制
错误
复制成功
 上一章节
获取运费模板
销售
下一章节 
查询刊登结果
销售