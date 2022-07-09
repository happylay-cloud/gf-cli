package gen

const templateDaoModelIndexContent = `
// ==========================================================================
//                这是由gf cli工具自动生成的。 根据需要填充此文件。
// ==========================================================================

package model

import (
	"{TplImportPrefix}/model/internal"
)

// {TplTableNameCamelCase} 是{TplTableName}表的golang结构体。
type {TplTableNameCamelCase} internal.{TplTableNameCamelCase}

// 在下面填充你的想法。
`

const templateDaoModelInternalContent = `
// ==========================================================================
// 			      这是由gf cli工具自动生成的。不要手动编辑此文件。
// ==========================================================================

package internal

{TplPackageImports}

// {TplTableNameCamelCase} 是{TplTableName}表的golang结构体。
{TplStructDefine}
`
