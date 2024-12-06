const fs = require('fs')
const path = require('path')
// 获取命令行参数
const args = process.argv.slice(2)


// 接收输入
var askForInput = null // 默认为空
process.stdin.resume()
process.stdin.on('data', (data) => {
    if (!askForInput) return;
    let text = data.toString().trim().toLocaleUpperCase()

    // 触发
    askForInput && askForInput(text)
    askForInput = null
})



function isReplace(filePath) {
    return new Promise((resolve, reject) => {
        console.log(`${filePath} \n已存在，是否替换？(y/n)`)
        askForInput = (text) => {
            if (text === 'Y') {
                resolve(true)
            } else {
                resolve(false)
            }
        }
    })
}


(async function () {

    // node .\scripts\gen.js ctx manage/user/loginInfo
    if (args[0] === 'ctx') {
        let ctxPath = args[1]
        let pathList = ctxPath.split('/')
        // 操作名称 驼峰转下划线
        let name = pathList.pop()
        let _name = name.replace(/([A-Z])/g, '_$1').toLowerCase()
        // 控制器名称
        let ctxName = [...pathList].map(e => e.charAt(0).toUpperCase() + e.slice(1)).join('')
        ctxName += name.charAt(0).toUpperCase() + name.slice(1)
        console.log(ctxName)

        let filePath = path.join(__dirname, '../', `internal/controller/${pathList.join('_')}__${_name}.go`)
        // let replace = await isReplace(filePath)
        if (fs.existsSync(filePath)) {
            let replace = await isReplace(filePath)
            if (!replace) {
                return
            }
        }

        let template = fs.readFileSync(path.join(__dirname, 'template/ctx.tmp'), 'utf8')
        let content = template.replace(/{{name}}/g, ctxName)
        fs.writeFileSync(filePath, content)

        // 路由名称
        let ___name = _name.replace(/_/g, '-')
        let routeName = ['/api', ...pathList, ___name].join('/')
        let methodName = (args[2] || 'GET').toLocaleUpperCase()
        let routeTemplate = fs.readFileSync(path.join(__dirname, 'template/router.tmp'), 'utf8')
        let routeContent = routeTemplate.replace(/{{path}}/g, routeName)
            .replace(/{{method}}/g, methodName)
            .replace(/{{name}}/g, ctxName)


        let routePath = path.join(__dirname, '../', `internal/router/${pathList.join('_')}__${_name}.go`)
        if (fs.existsSync(routePath)) {
            let replace = await isReplace(routePath)
            if (!replace) {
                return
            }
        }

        fs.writeFileSync(routePath, routeContent)
    }


    process.exit(0)
})()



// internal/controller/${模块}/${控制器}/${方法}.go


// manage_user__login.go
