const path = require('path');
module.exports = {
    publicPath: process.env.NODE_ENV == 'production' ? '/' : '/',
    devServer: {
        port: 1998, //mở cộng chạy mới
        open: true, //tự động mở trình duyệt khi build xong
    },

}