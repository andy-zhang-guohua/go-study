package sys_login_log

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/models/system"

	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/apis"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/tools"

	"net/http"
)

type SysLoginLog struct {
	apis.Api
}

func (e *SysLoginLog) GetSysLoginLogList(c *gin.Context) {
	msgID := tools.GenerateMsgIDFromContext(c)
	d := new(dto.SysLoginLogSearch)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	req := d.Generate()

	//查询列表
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	list := make([]system.SysLoginLog, 0)
	var count int64
	serviceStudent := service.SysLoginLog{}
	serviceStudent.MsgID = msgID
	serviceStudent.Orm = db
	err = serviceStudent.GetSysLoginLogPage(req, &list, &count)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.PageOK(c, list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

func (e *SysLoginLog) GetSysLoginLog(c *gin.Context) {
	control := new(dto.SysLoginLogById)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	//查看详情
	req := control.Generate()
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object system.SysLoginLog

	serviceSysLoginLog := service.SysLoginLog{}
	serviceSysLoginLog.MsgID = msgID
	serviceSysLoginLog.Orm = db
	err = serviceSysLoginLog.GetSysLoginLog(req, &object)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.OK(c, object, "查看成功")
}

func (e *SysLoginLog) InsertSysLoginLog(c *gin.Context) {
	control := new(dto.SysLoginLogControl)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	//新增操作
	req := control.Generate()
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object common.ActiveRecord
	object, err = req.GenerateM()
	if err != nil {
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	// 设置创建人
	object.SetCreateBy(tools.GetUserIdUint(c))

	serviceSysLoginLog := service.SysLoginLog{}
	serviceSysLoginLog.Orm = db
	serviceSysLoginLog.MsgID = msgID
	err = serviceSysLoginLog.InsertSysLoginLog(object)
	if err != nil {
		log.Error(err)
		e.Error(c, http.StatusInternalServerError, err, "创建失败")
		return
	}

	e.OK(c, object.GetId(), "创建成功")
}

func (e *SysLoginLog) UpdateSysLoginLog(c *gin.Context) {
	control := new(dto.SysLoginLogControl)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	req := control.Generate()
	//更新操作
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object common.ActiveRecord
	object, err = req.GenerateM()
	if err != nil {
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	object.SetUpdateBy(tools.GetUserIdUint(c))

	serviceSysLoginLog := service.SysLoginLog{}
	serviceSysLoginLog.Orm = db
	serviceSysLoginLog.MsgID = msgID
	err = serviceSysLoginLog.UpdateSysLoginLog(object)
	if err != nil {
		log.Error(err)
		return
	}
	e.OK(c, object.GetId(), "更新成功")
}

func (e *SysLoginLog) DeleteSysLoginLog(c *gin.Context) {
	control := new(dto.SysLoginLogById)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	//删除操作
	req := control.Generate()
	err = req.Bind(c)
	if err != nil {
		log.Errorf("MsgID[%s] Bind error: %s", msgID, err)
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object common.ActiveRecord
	object, err = req.GenerateM()
	if err != nil {
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}

	// 设置编辑人
	object.SetUpdateBy(tools.GetUserIdUint(c))

	serviceSysLoginLog := service.SysLoginLog{}
	serviceSysLoginLog.Orm = db
	serviceSysLoginLog.MsgID = msgID
	err = serviceSysLoginLog.RemoveSysLoginLog(req, object)
	if err != nil {
		log.Error(err)
		return
	}
	e.OK(c, object.GetId(), "删除成功")
}
