import { DataTypes, QueryInterface } from "sequelize";

module.exports = {
  async up (queryInterface: QueryInterface) {
    await queryInterface.addColumn('rooms', 'price', {
      type: DataTypes.INTEGER,
      allowNull: false,
    });
  },

  async down (queryInterface: QueryInterface) {
    await queryInterface.removeColumn('rooms', 'price');
  }
};
