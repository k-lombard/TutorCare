
const CompressionPlugin = require(`compression-webpack-plugin`);
const BrotliPlugin = require(`brotli-webpack-plugin`);
const path = require(`path`);
module.exports = {
    plugins:[
        new BrotliPlugin({
            asset: '[fileWithoutExt].[ext].br',
            test: /\.(js|scss|css|html|svg)$/
        })
    ],
}
