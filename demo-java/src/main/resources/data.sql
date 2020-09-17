insert into MOTORISATION(id, acceleration, autonomy) VALUES
(0,5.6,409),(1,4.6,560),(2,3.4,530);

insert into MODEL_TESLA(id, description, name, photo_link, type) VALUES
(0, 'Familiale routière', 'Model 3', '/resources/model3.jpeg', 'Berline');

insert into TESLA (id, price, model_tesla_id, motorisation_id) VALUES
(0, 49600, 0, 0), (1, 57800, 0, 1), (2, 64890, 0, 2);
