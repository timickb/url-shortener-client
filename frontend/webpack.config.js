const path = require('path')
const HTMLWebpackPlugin = require('html-webpack-plugin')
const { CleanWebpackPlugin } = require('clean-webpack-plugin')

module.exports = {
    context: path.resolve(__dirname, 'src'),
    mode: 'development',
    entry: './index.js',

    output: {
        filename: '[name].[hash].js',
        path: path.resolve(__dirname, 'dist')
    },

    resolve: {
        extensions: ['.js'],
        alias: {
            '@': path.resolve(__dirname, 'src')
        },
    },

    plugins: [
        new HTMLWebpackPlugin({template: './index.html'}),
        new CleanWebpackPlugin()
    ],

    devServer: {
        port: 9000
    },

    module: {
        rules: [
            { test: /\.css$/, use: ['style-loader', 'css-loader'] },
            { test: /\.(png|jp(e*)g|gif|svg)$/, use: ['file-loader'] },
        ]
    }
}