import express from 'express';
import request from 'supertest';
import * as hotelService from '../../src/services/hotel.service';
import router from '../../src/routers/v1/index.router';


jest.mock('../../src/services/hotel.service');

describe('Hotel Controller', () => {
  const app = express().use(express.json()).use('/api/v1', router);

  it('POST /api/v1/hotels return 201 with created hotel', async () => {
    (hotelService.createHotelService as jest.Mock).mockResolvedValue({ id: 1, name: 'Hotel A'});

    const res = await request(app)
      .post('/api/v1/hotels')
      .send({
        name: 'Hotel A',
        address: '123 Main St',
        location: 'Delhi',
        rating: 4.5,
        ratingCount: 10,
      });

    expect(res.status).toBe(201);
    expect(res.body?.data).toMatchObject({ id: 1, name: 'Hotel A'});
    expect(hotelService.createHotelService).toHaveBeenCalledWith(expect.objectContaining({ name: 'Hotel A'}));
  });

  it('handles service errors -> 400/500', async () => {
    (hotelService.createHotelService as jest.Mock).mockRejectedValue(new Error('Bad'));

    const res = await request(app).post('/api/v1/hotels').send({ name: 'X'});

    // assert your error middleware mapping
    expect(res.status).toBeGreaterThanOrEqual(400);
  });
});