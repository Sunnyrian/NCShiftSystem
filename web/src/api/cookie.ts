  let cookie = {
    /**
   * 设置cookie方法
   * @param key{string} cookie名称
   * @param val{string} cookie值
   * @param hours{number} 过期时间（小时）
   */
      set: function (key: string, val: string, hours: number) {
      let date = new Date(); //获取当前时间
      let expiresHours = hours;  //将date设置为n小时以后的时间
      date.setTime(date.getTime() + expiresHours * 24 * 3600 * 1000); //格式化为cookie识别的时间
      document.cookie = key + "=" + val + ";expires=" + date.toUTCString();  //设置cookie
    },
    /**
     * 移除cookie方法
     * @param key{string} cookie名称
     */
    remove: function (key: string) {
      document.cookie = 'token=;expires=Thu, 01-Jan-1970 00:00:01 GMT';
    },
    get: function (key:string) {//获取cookie方法
      /*获取cookie参数*/
      let getCookie = document.cookie.replace(/[ ]/g, "");  
      //获取cookie，并且将获得的cookie格式化，去掉空格字符
      let arrCookie = getCookie.split(";")  //将获得的cookie以"分号"为标识 将cookie保存到arrCookie的数组中
      let tips = "";  //声明变量tips
      for (let i = 0; i < arrCookie.length; i++) {   //使用for循环查找cookie中的tips变量
        let arr = arrCookie[i].split("=");   //将单条cookie用"等号"为标识，将单条cookie保存为arr数组
        if (key === arr[0]) {  //匹配变量名称，其中arr[0]是指的cookie名称，如果该条变量为tips则执行判断语句中的赋值操作
          tips = arr[1];   //将cookie的值赋给变量tips
          break;   //终止for循环遍历
        }
      }
      return tips;
    }
  }

  export default cookie