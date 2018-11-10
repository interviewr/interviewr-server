-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE person (
  id SERIAL PRIMARY KEY,
  login VARCHAR(20) NOT NULL,
  name VARCHAR(35) NOT NULL,
  email VARCHAR(35) UNIQUE NOT NULL,
  location VARCHAR(50),
  bio VARCHAR,
  avatarUrl VARCHAR
);

CREATE TABLE organization (
  id SERIAL PRIMARY KEY,
  name VARCHAR(35) NOT NULL,
  email VARCHAR(35) UNIQUE NOT NULL,
  description VARCHAR,
  location VARCHAR(50),
  avatarUrl VARCHAR
);

-- TODO:
-- https://laracasts.com/discuss/channels/eloquent/database-schema-for-users-belonging-to-groups-with-roles
CREATE TABLE membership (
  organization_id INT REFERENCES organization(id),
  person_id INT REFERENCES person(id),
  role VARCHAR CHECK (role IN ('owner', 'member')),
  CONSTRAINT membership_id PRIMARY KEY (organization_id, person_id)
);

CREATE TABLE vacancy (
  id SERIAL PRIMARY KEY,
  title VARCHAR(30) NOT NULL,
  description VARCHAR NOT NULL,
  location VARCHAR(50),
  organization_id INT REFERENCES organization(id)
);

-- PERSON:
INSERT INTO person (id, login, name, email, location, bio, avatarUrl) VALUES (1, 'lgiametti0', 'Luce Giametti', 'lgiametti0@yellowbook.com', 'Minnesota', 'In hac habitasse platea dictumst.', 'https://robohash.org/nemohicminima.bmp?size=50x50&set=set1');
INSERT INTO person (id, login, name, email, location, bio, avatarUrl) VALUES (2, 'wlowseley1', 'Ward Lowseley', 'wlowseley1@cdc.gov', 'Texas', 'Pellentesque eget nunc. Donec quis orci eget orci vehicula condimentum. Curabitur in libero ut massa volutpat convallis. Morbi odio odio, elementum eu, interdum eu, tincidunt in, leo. Maecenas pulvinar lobortis est. Phasellus sit amet erat. Nulla tempus. Vivamus in felis eu sapien cursus vestibulum. Proin eu mi.', 'https://robohash.org/autemerrorquia.png?size=50x50&set=set1');
INSERT INTO person (id, login, name, email, location, bio, avatarUrl) VALUES (3, 'mremirez2', 'Myrah Remirez', 'mremirez2@go.com', 'New York', 'Maecenas rhoncus aliquam lacus.', 'https://robohash.org/etquaeratquidem.png?size=50x50&set=set1');
INSERT INTO person (id, login, name, email, location, bio, avatarUrl) VALUES (4, 'mdebling3', 'Marin Debling', 'mdebling3@cyberchimps.com', 'Colorado', 'Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Duis faucibus accumsan odio. Curabitur convallis. Duis consequat dui nec nisi volutpat eleifend. Donec ut dolor. Morbi vel lectus in quam fringilla rhoncus. Mauris enim leo, rhoncus sed, vestibulum sit amet, cursus id, turpis. Integer aliquet, massa id lobortis convallis, tortor risus dapibus augue, vel accumsan tellus nisi eu orci.', 'https://robohash.org/voluptatumquaealias.png?size=50x50&set=set1');
INSERT INTO person (id, login, name, email, location, bio, avatarUrl) VALUES (5, 'hdaubney4', 'Heloise D''Aubney', 'hdaubney4@ucla.edu', 'California', 'Proin eu mi. Nulla ac enim. In tempor, turpis nec euismod scelerisque, quam turpis adipiscing lorem, vitae mattis nibh ligula nec sem. Duis aliquam convallis nunc. Proin at turpis a pede posuere nonummy. Integer non velit. Donec diam neque, vestibulum eget, vulputate ut, ultrices vel, augue.', 'https://robohash.org/exercitationemundeofficiis.bmp?size=50x50&set=set1');
INSERT INTO person (id, login, name, email, location, bio, avatarUrl) VALUES (6, 'bpellingar5', 'Barbara-anne Pellingar', 'bpellingar5@stanford.edu', 'Texas', 'Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Mauris viverra diam vitae quam. Suspendisse potenti. Nullam porttitor lacus at turpis. Donec posuere metus vitae ipsum. Aliquam non mauris. Morbi non lectus.', 'https://robohash.org/perspiciatisaccusantiumautem.bmp?size=50x50&set=set1');
INSERT INTO person (id, login, name, email, location, bio, avatarUrl) VALUES (7, 'dchesshyre6', 'Dermot Chesshyre', 'dchesshyre6@blogtalkradio.com', 'California', 'Etiam faucibus cursus urna. Ut tellus.', 'https://robohash.org/omnisveritatisvoluptas.png?size=50x50&set=set1');
INSERT INTO person (id, login, name, email, location, bio, avatarUrl) VALUES (8, 'mmapledorum7', 'Mariellen Mapledorum', 'mmapledorum7@macromedia.com', 'Virginia', 'Duis mattis egestas metus. Aenean fermentum. Donec ut mauris eget massa tempor convallis. Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh. Quisque id justo sit amet sapien dignissim vestibulum.', 'https://robohash.org/minimainratione.png?size=50x50&set=set1');
INSERT INTO person (id, login, name, email, location, bio, avatarUrl) VALUES (9, 'vdilliway8', 'Vale Dilliway', 'vdilliway8@51.la', 'Georgia', 'Praesent id massa id nisl venenatis lacinia.', 'https://robohash.org/etsimiliquemolestiae.jpg?size=50x50&set=set1');
INSERT INTO person (id, login, name, email, location, bio, avatarUrl) VALUES (10, 'tslocumb9', 'Tiena Slocumb', 'tslocumb9@theguardian.com', 'Texas', 'Aliquam quis turpis eget elit sodales scelerisque. Mauris sit amet eros. Suspendisse accumsan tortor quis turpis. Sed ante. Vivamus tortor. Duis mattis egestas metus. Aenean fermentum.', 'https://robohash.org/delenitifacilisvel.bmp?size=50x50&set=set1');

-- ORGANIZATION:
INSERT INTO organization (id, name, email, description, location, avatarUrl) VALUES (1, 'Mynte', 'esevern0@e-recht24.de', 'Mauris lacinia sapien quis libero. Nullam sit amet turpis elementum ligula vehicula consequat. Morbi a ipsum. Integer a nibh. In quis justo. Maecenas rhoncus aliquam lacus. Morbi quis tortor id nulla ultrices aliquet. Maecenas leo odio, condimentum id, luctus nec, molestie sed, justo. Pellentesque viverra pede ac diam. Cras pellentesque volutpat dui.', 'Texas', 'https://robohash.org/debitisquasisit.bmp?size=50x50&set=set1');
INSERT INTO organization (id, name, email, description, location, avatarUrl) VALUES (2, 'Edgetag', 'tbeasant1@seesaa.net', 'Nulla suscipit ligula in lacus. Curabitur at ipsum ac tellus semper interdum. Mauris ullamcorper purus sit amet nulla. Quisque arcu libero, rutrum ac, lobortis vel, dapibus at, diam. Nam tristique tortor eu pede.', 'Texas', 'https://robohash.org/rerumquiaitaque.jpg?size=50x50&set=set1');
INSERT INTO organization (id, name, email, description, location, avatarUrl) VALUES (3, 'Tazz', 'cmorhall2@seesaa.net', 'Sed sagittis. Nam congue, risus semper porta volutpat, quam pede lobortis ligula, sit amet eleifend pede libero quis orci. Nullam molestie nibh in lectus. Pellentesque at nulla. Suspendisse potenti.', 'District of Columbia', 'https://robohash.org/nonautemcum.bmp?size=50x50&set=set1');
INSERT INTO organization (id, name, email, description, location, avatarUrl) VALUES (4, 'Youfeed', 'jantal3@instagram.com', 'Nam nulla. Integer pede justo, lacinia eget, tincidunt eget, tempus vel, pede. Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus. In est risus, auctor sed, tristique in, tempus sit amet, sem. Fusce consequat. Nulla nisl. Nunc nisl. Duis bibendum, felis sed interdum venenatis, turpis enim blandit mi, in porttitor pede justo eu massa.', 'Michigan', 'https://robohash.org/aofficiisculpa.png?size=50x50&set=set1');
INSERT INTO organization (id, name, email, description, location, avatarUrl) VALUES (5, 'Innojam', 'stozer4@mapy.cz', 'Donec posuere metus vitae ipsum. Aliquam non mauris. Morbi non lectus. Aliquam sit amet diam in magna bibendum imperdiet.', 'District of Columbia', 'https://robohash.org/repellendusquiadebitis.jpg?size=50x50&set=set1');
INSERT INTO organization (id, name, email, description, location, avatarUrl) VALUES (6, 'Leenti', 'ckochlin5@cbslocal.com', 'Nam dui. Proin leo odio, porttitor id, consequat in, consequat ut, nulla. Sed accumsan felis. Ut at dolor quis odio consequat varius. Integer ac leo. Pellentesque ultrices mattis odio. Donec vitae nisi. Nam ultrices, libero non mattis pulvinar, nulla pede ullamcorper augue, a suscipit nulla elit ac nulla.', 'California', 'https://robohash.org/numquamcupiditatevoluptate.jpg?size=50x50&set=set1');
INSERT INTO organization (id, name, email, description, location, avatarUrl) VALUES (7, 'Muxo', 'ybothen6@columbia.edu', 'Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Donec pharetra, magna vestibulum aliquet ultrices, erat tortor sollicitudin mi, sit amet lobortis sapien sapien non mi. Integer ac neque. Duis bibendum. Morbi non quam nec dui luctus rutrum. Nulla tellus. In sagittis dui vel nisl.', 'District of Columbia', 'https://robohash.org/etetdistinctio.bmp?size=50x50&set=set1');
INSERT INTO organization (id, name, email, description, location, avatarUrl) VALUES (8, 'Tagchat', 'eduigenan7@histats.com', 'Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh. Quisque id justo sit amet sapien dignissim vestibulum. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Nulla dapibus dolor vel est.', 'Maryland', 'https://robohash.org/molestiaeperspiciatisincidunt.jpg?size=50x50&set=set1');
INSERT INTO organization (id, name, email, description, location, avatarUrl) VALUES (9, 'Devpulse', 'kdibaudi8@hugedomains.com', 'Mauris sit amet eros. Suspendisse accumsan tortor quis turpis. Sed ante. Vivamus tortor.', 'Florida', 'https://robohash.org/accusamusautqui.jpg?size=50x50&set=set1');
INSERT INTO organization (id, name, email, description, location, avatarUrl) VALUES (10, 'Fliptune', 'greadwood9@economist.com', 'Integer ac neque. Duis bibendum. Morbi non quam nec dui luctus rutrum. Nulla tellus. In sagittis dui vel nisl. Duis ac nibh.', 'Pennsylvania', 'https://robohash.org/inciduntautmolestias.png?size=50x50&set=set1');

-- MEMBERSHIP:
INSERT INTO membership (organization_id, person_id, role) VALUES (1, 1, 'owner');
INSERT INTO membership (organization_id, person_id, role) VALUES (1, 2, 'member');
INSERT INTO membership (organization_id, person_id, role) VALUES (1, 3, 'member');
INSERT INTO membership (organization_id, person_id, role) VALUES (2, 4, 'owner');
INSERT INTO membership (organization_id, person_id, role) VALUES (2, 5, 'member');
INSERT INTO membership (organization_id, person_id, role) VALUES (8, 6, 'owner');
INSERT INTO membership (organization_id, person_id, role) VALUES (4, 7, 'owner');
INSERT INTO membership (organization_id, person_id, role) VALUES (4, 8, 'member');
INSERT INTO membership (organization_id, person_id, role) VALUES (10, 9, 'owner');
INSERT INTO membership (organization_id, person_id, role) VALUES (6, 10, 'owner');

-- VACANCY:
INSERT INTO vacancy (id, title, description, location, organization_id) VALUES (1, 'Professor', 'Sed accumsan felis. Ut at dolor quis odio consequat varius. Integer ac leo. Pellentesque ultrices mattis odio. Donec vitae nisi. Nam ultrices, libero non mattis pulvinar, nulla pede ullamcorper augue, a suscipit nulla elit ac nulla. Sed vel enim sit amet nunc viverra dapibus.', 'California', 1);
INSERT INTO vacancy (id, title, description, location, organization_id) VALUES (2, 'Senior Financial Analyst', 'Vestibulum ac est lacinia nisi venenatis tristique. Fusce congue, diam id ornare imperdiet, sapien urna pretium nisl, ut volutpat sapien arcu sed augue.', 'California', 1);
INSERT INTO vacancy (id, title, description, location, organization_id) VALUES (3, 'Dental Hygienist', 'Vestibulum rutrum rutrum neque. Aenean auctor gravida sem.', 'Pennsylvania', 8);
INSERT INTO vacancy (id, title, description, location, organization_id) VALUES (4, 'Human Resources Manager', 'Fusce posuere felis sed lacus. Morbi sem mauris, laoreet ut, rhoncus aliquet, pulvinar sed, nisl.', 'Florida', 8);
INSERT INTO vacancy (id, title, description, location, organization_id) VALUES (5, 'Staff Scientist', 'Aliquam non mauris. Morbi non lectus.', 'Nebraska', 4);
INSERT INTO vacancy (id, title, description, location, organization_id) VALUES (6, 'Research Nurse', 'In hac habitasse platea dictumst. Maecenas ut massa quis augue luctus tincidunt. Nulla mollis molestie lorem. Quisque ut erat. Curabitur gravida nisi at nibh. In hac habitasse platea dictumst. Aliquam augue quam, sollicitudin vitae, consectetuer eget, rutrum at, lorem. Integer tincidunt ante vel ipsum.', 'Virginia', 4);
INSERT INTO vacancy (id, title, description, location, organization_id) VALUES (7, 'Statistician II', 'In hac habitasse platea dictumst. Etiam faucibus cursus urna. Ut tellus. Nulla ut erat id mauris vulputate elementum. Nullam varius. Nulla facilisi.', 'California', 10);
INSERT INTO vacancy (id, title, description, location, organization_id) VALUES (8, 'Environmental Specialist', 'Nullam sit amet turpis elementum ligula vehicula consequat. Morbi a ipsum. Integer a nibh. In quis justo.', 'Connecticut', 6);
INSERT INTO vacancy (id, title, description, location, organization_id) VALUES (9, 'Account Representative I', 'Maecenas pulvinar lobortis est. Phasellus sit amet erat. Nulla tempus. Vivamus in felis eu sapien cursus vestibulum. Proin eu mi. Nulla ac enim. In tempor, turpis nec euismod scelerisque, quam turpis adipiscing lorem, vitae mattis nibh ligula nec sem.', 'Missouri', 6);
INSERT INTO vacancy (id, title, description, location, organization_id) VALUES (10, 'Account Representative III', 'Duis at velit eu est congue elementum. In hac habitasse platea dictumst. Morbi vestibulum, velit id pretium iaculis, diam erat fermentum justo, nec condimentum neque sapien placerat ante. Nulla justo. Aliquam quis turpis eget elit sodales scelerisque. Mauris sit amet eros. Suspendisse accumsan tortor quis turpis. Sed ante.', 'Nevada', 1);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE membership;
DROP TABLE person;
DROP TABLE organization;
