const path = require('path');
const HtmlWebpackPlugin = require("html-webpack-plugin");

module.exports = {
    mode: 'development',
    entry: './app/app.jsx',
    devServer: {
        contentBase: [
            path.join(__dirname, 'dist'),
            path.join(__dirname, 'assets'),
        ],
        compress: true,
        hot: true,
        disableHostCheck: true,
        host: '0.0.0.0',
        port: 8080
    },
    module: {
        rules: [   //загрузчик для jsx
            {
                test: /\.jsx?$/, // определяем тип файлов
                exclude: /(node_modules)/,  // исключаем из обработки папку node_modules
                loader: "babel-loader",   // определяем загрузчик
                options: {
                    presets: ["@babel/preset-env", "@babel/preset-react"]    // используемые плагины
                }
            },
            {
                test: /\.js$/,
                use: {
                    loader: "babel-loader"
                }
            },
            {
                test: /.css$/,
                use: [
                    'style-loader',
                    'css-loader',
                ]
            },
            {
                test: /\.s[ac]ss$/i,
                use: [
                    'style-loader',
                    'sass-loader',
                ]
            },
            {
                test: /\.(png|svg|jpg|jpeg|gif)$/i,
                use: [
                    {
                        loader: 'file-loader',
                        options: {
                            name: '[name].[contenthash].[ext]',
                            esModule: false, // <- here
                        }
                    },
                ],
            },
        ],
    },
    plugins: [
        new HtmlWebpackPlugin({
            template: "./index.html"
        })
    ]
}
