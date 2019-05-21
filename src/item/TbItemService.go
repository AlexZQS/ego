package item

import (
	"ego/src/common"
	"ego/src/item/cat"
	"ego/src/item/desc"
	"fmt"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func showItemService(page, rows int) (e *common.DataGrid) {
	items := selByPageDao(rows, page)
	if items != nil {
		itemChildren := make([]TbItemChild, 0)
		for i := 0; i < len(items); i++ {
			var itemChild TbItemChild
			itemChild.Id = items[i].Id
			itemChild.Barcode = items[i].Barcode
			itemChild.Updated = items[i].Updated
			itemChild.Created = items[i].Created
			itemChild.Status = items[i].Status
			//itemChild.Cid = items[i].Cid
			//itemChild.Image = items[i].Image
			itemChild.Num = items[i].Num
			itemChild.Price = items[i].Price
			itemChild.Image = items[i].Image
			itemChild.SellPoint = items[i].SellPoint
			itemChild.Title = items[i].Title

			fmt.Printf(" id = %d", items[i].Cid)
			itemChild.CategoryName = cat.ShowCatByIdService(items[i].Cid).Name
			fmt.Printf(" CategoryName = %s\n", itemChild.CategoryName)
			itemChildren = append(itemChildren, itemChild)

		}
		e = new(common.DataGrid)
		e.Row = itemChildren //当前页显示的数据
		e.Total = selCount()
		return
	}
	return nil
}

//删除商品
func delByIdsService(ids string) (e common.EgoResult) {
	count := updStatusByIdsDao(strings.Split(ids, ","), 3)
	if count > 0 {
		e.Status = 200

	}
	return
}

//商品上架
func instockService(ids string) (e common.EgoResult) {
	count := updStatusByIdsDao(strings.Split(ids, ","), 1)
	if count > 0 {
		e.Status = 200
	}
	return
}

//商品下架
func offstockService(ids string) (e common.EgoResult) {
	count := updStatusByIdsDao(strings.Split(ids, ","), 2)
	if count > 0 {
		e.Status = 200
	}
	return
}

//图片上传
func imageUploadService(f multipart.File, h *multipart.FileHeader) map[string]interface{} {
	m := make(map[string]interface{})
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		m["error"] = 1
		m["message"] = "上传失败,服务器错误"
		return m
	}

	//纳秒时间戳 + 随机数 + 扩展名
	rand.Seed(time.Now().UnixNano()) //种子
	fileName := "static/images/" +
		strconv.FormatInt(time.Now().UnixNano(), 12) +
		strconv.Itoa(rand.Intn(1000)) +
		h.Filename[strings.LastIndex(h.Filename, "."):]
	fmt.Printf("FileName %s", fileName)
	err = ioutil.WriteFile(fileName, bytes, 0777)
	if err != nil {
		m["error"] = 1
		m["message"] = "上传失败,保存图片时错误"
		return m
	}
	m["error"] = 0
	m["url"] = common.CurrPath + fileName
	return m
}

//商品新增
func insertService(form url.Values) (e common.EgoResult) {
	var t TbItem

	cid, _ := strconv.Atoi(form["Cid"][0])
	t.Cid = cid
	t.Title = form["Title"][0]
	t.SellPoint = form["SellPoint"][0]

	price, _ := strconv.Atoi(form["Price"][0])
	t.Price = price

	num, _ := strconv.Atoi(form["Num"][0])
	t.Num = num
	t.Image = form["Image"][0]

	t.Status = 1
	timeFormat := time.Now().Format("2006-01-02 15:04:05")
	t.Created = timeFormat

	t.Updated = timeFormat

	id := common.GenId()
	t.Id = id

	//商品表新增
	count := insertItemDao(t)
	if count > 0 {
		var tbItemDesc desc.TbItemDesc
		tbItemDesc.ItemId = id
		tbItemDesc.Created = timeFormat
		tbItemDesc.Update = timeFormat
		tbItemDesc.ItemDesc = form["Desc"][0]
		countDesc := desc.Insert(tbItemDesc)
		if countDesc > 0 {
			e.Status = 200
		} else {
			//删除商品中数据
			delById(id)
		}
	}
	return
}

//修改页面显示信息
func showItemDescCatService(id int) TbItemDescChild {
	item := selByIdDao(id)
	var c TbItemDescChild
	c.Id = item.Id
	c.Updated = item.Updated
	c.Created = item.Created
	c.Barcode = item.Barcode
	c.Cid = item.Cid
	c.Title = item.Title
	c.SellPoint = item.SellPoint
	c.Price = item.Price
	c.Image = item.Image
	c.Status = item.Status
	c.Num = item.Num
	//商品类目
	c.CategoryName = cat.ShowCatByIdService(c.Cid).Name
	//商品描述
	c.Desc = desc.SelByIdService(c.Id).ItemDesc
	return c
}

//更新
func update(values url.Values) (e common.EgoResult) {
	common.OpenConnWithTx()
	var t TbItem
	id, _ := strconv.Atoi(values["Id"][0])
	t.Id = int64(id)

	cid, _ := strconv.Atoi(values["Cid"][0])
	t.Cid = cid

	t.Title = values["Title"][0]

	price, _ := strconv.Atoi(values["Price"][0])
	t.Price = price

	num, _ := strconv.Atoi(values["Num"][0])
	t.Num = num

	t.Image = values["Image"][0]

	status, _ := strconv.Atoi(values["Status"][0])
	t.Status = int8(status)

	t.SellPoint = values["SellPoint"][0]

	date := time.Now().Format("2006-01-02 15:04:05")
	t.Updated = date

	count := updateItemByIdWithTx(t)
	if count > 0 {
		var itemDesc desc.TbItemDesc
		itemDesc.ItemId = int64(id)
		itemDesc.ItemDesc = values["Desc"][0]
		itemDesc.Update = date
		count = desc.UpdateDescByIdWithTxDao(itemDesc)
		if count > 0 || "" == itemDesc.ItemDesc {
			common.CloseConnWithTx(true)
			e.Status = 200
		}
	}
	common.CloseConnWithTx(false)
	return

}
