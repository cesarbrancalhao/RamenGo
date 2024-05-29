INSERT INTO broth (id, image_inactive, image_active, name, description, price)
VALUES ('1',
    'https://tech.redventures.com.br/icons/salt/inactive.svg',
    'https://tech.redventures.com.br/icons/salt/active.svg',
    'Salt',
    'Simple like the seawater, nothing more',
    10);

INSERT INTO broth (id, image_inactive, image_active, name, description, price)
VALUES ('2',
    'https://tech.redventures.com.br/icons/shoyu/inactive.svg',
    'https://tech.redventures.com.br/icons/shoyu/active.svg',
    'Shoyu',
    'The good old and traditional soy sauce',
    10);

INSERT INTO broth (id, image_inactive, image_active, name, description, price)
VALUES ('3',
    'https://tech.redventures.com.br/icons/miso/inactive.svg',
    'https://tech.redventures.com.br/icons/miso/active.svg',
    'Miso',
    'Paste made of fermented soybeans',
    12);

INSERT INTO protein (id, image_inactive, image_active, name, description, price)
VALUES ('1', 
    'https://tech.redventures.com.br/icons/pork/inactive.svg',
    'https://tech.redventures.com.br/icons/pork/active.svg',
    'Chasu',
    'A sliced flavourful pork meat with a selection of season vegetables.',
    10);

INSERT INTO protein (id, image_inactive, image_active, name, description, price)
VALUES ('2',
    'https://tech.redventures.com.br/icons/yasai/inactive.svg',
    'https://tech.redventures.com.br/icons/yasai/active.svg',
    'Yasai Vegetarian',
    'A delicious vegetarian lamen with a selection of season vegetables.',
    10);

INSERT INTO protein (id, image_inactive, image_active, name, description, price)
VALUES ('3',
    'https://tech.redventures.com.br/icons/chicken/inactive.svg',
    'https://tech.redventures.com.br/icons/chicken/active.svg', 
    'Karaague',
    'Three units of fried chicken, moyashi, ajitama egg and other vegetables.',
    12);

INSERT INTO order_response (id, description, image)
VALUES (1, 'Salt and Chasu Ramen', 'https://tech.redventures.com.br/icons/ramen/ramenChasu.png');

INSERT INTO order_response (id, description, image)
VALUES (2, 'Salt and Yasai Vegetarian Ramen', 'https://tech.redventures.com.br/icons/ramen/ramenYasai Vegetarian.png');

INSERT INTO order_response (id, description, image)
VALUES (3, 'Salt and Karaague Ramen', 'https://tech.redventures.com.br/icons/ramen/ramenKaraague.png');

INSERT INTO order_response (id, description, image)
VALUES (4, 'Shoyu and Chasu Ramen', 'https://tech.redventures.com.br/icons/ramen/ramenChasu.png');

INSERT INTO order_response (id, description, image)
VALUES (5, 'Shoyu and Yasai Vegetarian Ramen', 'https://tech.redventures.com.br/icons/ramen/ramenYasai Vegetarian.png');

INSERT INTO order_response (id, description, image)
VALUES (6, 'Shoyu and Karaague Ramen', 'https://tech.redventures.com.br/icons/ramen/ramenKaraague.png');

INSERT INTO order_response (id, description, image)
VALUES (7, 'Miso and Chasu Ramen', 'https://tech.redventures.com.br/icons/ramen/ramenChasu.png');

INSERT INTO order_response (id, description, image)
VALUES (8, 'Miso and Yasai Vegetarian Ramen', 'https://tech.redventures.com.br/icons/ramen/ramenYasai Vegetarian.png');

INSERT INTO order_response (id, description, image)
VALUES (9, 'Miso and Karaague Ramen', 'https://tech.redventures.com.br/icons/ramen/ramenKaraague.png');

-- This is all the data provided, if you intend to add more I suggest to go to schema and turn ids into auto-increment