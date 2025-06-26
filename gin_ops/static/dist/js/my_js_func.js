//判断字符串是否为邮箱
function isValidEmail(email) {
  const regex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
  return regex.test(email);
}

function save(key, data) {
  localStorage.setItem(key, data);
}
function load(key) {
  return localStorage.getItem(key);
}
function dele(key) {
  localStorage.removeItem(key);
}

function save_json(key, data) {
  save(key, JSON.stringify(data));
}

function load_json(key) {
  var js_data = load(key);
  if (js_data) {
    return JSON.parse(js_data);
  } else {
    return null;
  }
}

function post_json(path, json, callback) {
  var host = "";
  var port = 0;
  var head_path = "/api/v1";
  //把cookie插入json
  var data = {};
  data["data"] = json;
  var cookie = load_json("cookie");
  if (cookie) {
    data["cookie"] = cookie;
  }
  var re_data = {};

  axios
    .post(head_path + path, data, {
      headers: {
        "Content-Type": "application/json",
      },
    })
    .then((response) => {
      //console.log(response)
      re_data["statusCode"] = response.status;
      //载入服务器返回的数据
      if (response.data) {
        re_data["data"] = response.data;
      }
      //自动保存服务器发送的cookie
      if (response.cookie) {
        if (response.cookie.Value == "") {
          dele("cookie");
        } else {
          save_json("cookie", response.cookie);
        }
      }
      callback(re_data);
    })
    .catch((error) => {
      re_data["statusCode"] = -1;
      re_data["error"]=error;
      callback(re_data);
    });
}
