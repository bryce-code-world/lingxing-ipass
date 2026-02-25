刊登流程说明
整体流程
选择要刊登的店铺和 Amazon 类目，确认 productType
获取指定 productType 的 JSON Schema（包含所有字段及校验规则）
本地处理 JSON Schema，填写商品资料
提交商品资料，等待亚马逊处理
查询刊登结果
步骤一：选择要刊登的店铺和 Amazon 类目，确认 productType
1. 请求【查询 Amazon 根分类】接口

查询 Amazon 根分类

接口地址：/basicOpen/openapi/publish/manage/categoryRoot
参数：要刊登的店铺 sid
返回：该站点的全部根分类及其 category_unique_id，用于查询子分类
2. 请求【查询 Amazon 子分类】接口

查询 Amazon 子分类

接口地址：/basicOpen/openapi/publish/manage/categoryChildren
参数：sid、category_unique_id
返回：该根分类的子分类及其 category_unique_id
流程说明：
如果子分类的 has_children=1，继续请求本接口，获取下一级子分类；
直到 has_children=0，该分类是最末级分类；
末级分类的：
productTypeOrigin 字段是商品类型，用于下一步获取刊登表单；
browseNodeAttributes 字段是节点位置，在第四步构建商品资料时需要用到，清注意留存。
美国站 返回 item_type_keyword，填到表单的同名字段（没有则填 productType 兜底）
非美国站 返回 recommended_browse_nodes，填到表单的同名字段
步骤二：获取指定 productType 的 JSON Schema

获取指定 productType 的 JSON Schema

接口地址：/basicOpen/openapi/publish/manage/getProductType
参数：marketplaceId、productTypeOrigin
返回：该商品类型两份表单
properties：站点语言
propertiesZh：中文版
步骤三：本地处理 JSON Schema，填写商品资料
每份刊登表单包含：
properties：所有字段，包括名称、定义、类型、限制、枚举值等
allOf：字段定义的判断条件
根据表单构建商品资料。已有多种开源库和工具可验证数据、呈现 UI 及生成代码。完整列表参见 JSON 架构工具。
特别注意：
刊登 FBM 配送商品需填写 merchant_shipping_group 字段（运费模板 id，非名称）。
获取运费模板

获取运费模板

接口地址：/basicOpen/openapi/publish/manage/getMerchantShippingGroup
参数：
seller_id、marketplace_id、product_type、flag
flag=0 返回已获取的运费模板（查询快）
flag=1 实时请求后台最新数据（查询慢，建议优先用 flag=0）
返回：数组，name 为运费模板名称，value 为运费模板 id（资料需填写模板 id）
步骤四：提交商品资料，等待亚马逊处理

提交商品资料

按 JSON Schema 格式生成商品 JSON，放在该 SKU 的 attributes 里
接口地址：/listing/publish/openapi/amazon/product/publish
支持批量提交：一次提交同个店铺的多个商品
不同 productType 的商品需重复步骤一~三，分别构建资料，然后批量提交
请求示例
{
  "store_id": xxxxx,
  "data": [
    {
      "sku": "xxxxxxx",
      "productType": "xxxxx",
      "attributes": {
        "item_name": [
          {
            "value": "MyBrand Carry-On Luggage",
            "language_tag": "en_US",
            "marketplace_id": "xxxxxxxxxxxxxxxx"
          }
        ],
        "brand": [
          {
            "value": "Generic",
            "language_tag": "en_US",
            "marketplace_id": "xxxxxxxxxxxxxxxx"
          }
        ],
        "supplier_declared_has_product_identifier_exemption": [
          {
            "value": true,
            "marketplace_id": "xxxxxxxxxxxxxxxx"
          }
        ],
        "item_type_keyword": [
          {
            "value": "xxxxxxxxxx",
            "marketplace_id": "xxxxxxxxxxxxxxxx"
          }
        ]
        // ... 更多属性
      }
    }
    // ... 更多 SKU
  ]
}
复制
错误
复制成功
步骤五：查询刊登结果

查询刊登结果

请求【查询刊登结果】接口
接口地址：/listing/publish/openapi/amazon/product/list
请求参数：
传入步骤四响应中的 record_unique_id，可查询该批次各商品的刊登结果；
也支持按 sku、store_id、operate_time 查询刊登结果，多个参数同时查询时为交集；
返回说明
status=0：表示该商品还在刊登中，请稍后重新轮询（亚马逊一般在 10 分钟内处理完成）
status=1：表示该商品刊登成功
status=2：表示该商品刊登失败，failure_reason 字段记录亚马逊的处理报错
需按照报错内容修改该商品资料后，重新请求【发布商品】接口
注意该批次已刊登成功的商品不要再重复提交
构建商品资料常见问题
Q1：没有品牌的商品，如何填写 brand？

A1：如商品无明确品牌，可在 brand 字段填写 "generic"。

⚠️ 某些情况下仍可能报错，具体请查看后台关于「与亚马逊无品牌商品政策相关的错误消息」的说明：帮助链接

Q2：没有商品编码的商品，如何填写？

A2：

externally_assigned_product_identifier 不填写
supplier_declared_has_product_identifier_exemption 的 value 传 true

注意：

必须先在亚马逊后台申请对应类目的 GTIN 豁免（流程详见后台说明：GTIN豁免申请）。
GTIN豁免按分类分别申请，部分分类不支持豁免。详见「各分类商品编码（全球贸易项目代码）要求」的说明：帮助链接
Q3：如何刊登变体商品？

A3：

父商品信息必须包含以下属性：
parentage_level：value 属性设为 parent
child_parent_sku_relationship：child_relationship_type 属性设为 variation，不传 parent_sku
variation_theme：name 属性设为适用的值
子商品信息需要具备：
parentage_level：value 属性设为 child
child_parent_sku_relationship：child_relationship_type 属性设为 variation，且 parent_sku 属性填写父商品的 SKU 编码
variation_theme：name 属性与父商品一致

在为子商品选择变体主题后，相应属性变为必填项
例如，variation_theme 设为 SIZE、COLOR 或 NUMBER_OF_ITEMS，则 shirt_size、color、number_of_items 均变为必填。

请根据所选变体主题，通过 ProductTypeDefinition 的条件要求（allOf）确定哪些属性需必填。

选择 variation_theme 时须参考 $lifecycle 下的 enumDeprecated。被标记为 enumDeprecated 的枚举值为已弃用值，请勿使用。

 上一章节
查询调价队列
销售
下一章节 
查询 Amazon 根分类
销售