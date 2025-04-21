import { dbConfig } from './index';

const config = {
  development: {
    username: dbConfig.DB_HOST,
    password: dbConfig.DB_PASSWORD,
    database: dbConfig.DB_NAME,
    host: dbConfig.DB_HOST,
    dialect: 'mysql',
  }
}

module.exports = config;