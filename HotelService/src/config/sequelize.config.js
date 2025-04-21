// This lie enables TypeScript support for the Sequelize configuration file
require('ts-node/register');
const config = require('./db.config');
module.exports = config;