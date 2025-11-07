CREATE TABLE users (
    id varchar(50) primary key,
    username varchar(255) not null,
    email varchar(255) not null unique,
    password varchar(255) not null,
    deposit_amount decimal(12, 2) default 0,
    role varchar(20) default 'user'
);

CREATE TABLE categories (
    id serial primary key,
    name varchar(255) not null,
    description text
);


CREATE TABLE equipments (
    id serial primary key,
    name varchar(255) not null,
    category_id int references categories(id) not null,
    description text,
    price_per_day decimal(12, 2) not null,
    price_per_week decimal(12, 2) not null,
    price_per_month decimal(12, 2) not null,
    price_per_year decimal(12, 2) not null,
    available boolean
);


CREATE TABLE rentals (
    id serial primary key,
    user_id varchar(50) references users(id) not null,
    equipment_id int references equipments(id) not null,
    rental_period varchar(20) not null,
    start_date timestamp,
    end_date timestamp,
    price decimal(12, 2) not null,
    status varchar(20) not null, --pending, active, cancelled, completed
    created_at timestamp default current_timestamp not null
);

CREATE TABLE payments (
    id serial primary key,
    user_id varchar(50) references users(id) not null,
    rental_id int references rentals(id) not null,
    amount decimal(12, 2) not null,
    payment_method varchar(20),
    status varchar(20) not null,
    created_at timestamp default current_timestamp
);

CREATE TABLE rental_histories (
    id serial primary key,
    rental_id int references rentals(id) not null,
    user_id varchar(50) references users(id) not null,
    status varchar(20) not null,
    created_at timestamp default current_timestamp
);

INSERT INTO categories (name, description) VALUES 
('Forklift', 'Forklift adalah alat berat yang digunakan untuk mengangkat dan memindahkan material berat dengan menggunakan garpu atau tine pada bagian depannya. Forklift sangat penting dalam operasi gudang dan pergudangan, memungkinkan pengangkutan barang dengan efisien dan aman.'),
('Crane', 'Crane adalah alat berat yang digunakan untuk mengangkat beban berat secara vertikal dan horisontal. Dengan kemampuan daya angkat yang tinggi, crane ideal untuk proyek-proyek konstruksi, pemasangan struktur, dan pemindahan material berat di ketinggian.'),
('Stoom', 'Stoom atau steamroller adalah alat berat yang digunakan untuk meratakan dan memadatkan permukaan tanah atau aspal. Stoom efektif dalam pembangunan jalan dan proyek infrastruktur lainnya, memastikan permukaan yang kuat dan tahan lama.'),
('Excavator', 'Excavator adalah alat berat serbaguna yang digunakan untuk pekerjaan penggalian, pemindahan tanah, dan konstruksi fondasi. Dengan lengan yang dapat berputar 360 derajat, excavator cocok untuk proyek-proyek konstruksi yang memerlukan presisi dan fleksibilitas.'),
('Bulldozer', 'Bulldozer adalah alat berat dengan pisau rata di bagian depannya, digunakan untuk meratakan dan membentuk tanah. Bulldozer sangat efektif dalam pekerjaan pembangunan jalan, land clearing, dan pengurukan tanah.'),
('Dozer Shovel', 'Dozer Shovel, juga dikenal sebagai loader shovel, adalah alat berat yang digunakan untuk memuat material seperti tanah, kerikil, atau batu. Cocok untuk proyek-proyek konstruksi dan pertambangan yang membutuhkan pemindahan material dengan cepat.'),
('Wheel Loader', 'Wheel Loader adalah alat berat dengan roda yang digunakan untuk memuat material ke truk atau conveyor. Wheel Loader memiliki kapasitas muat yang besar, membuatnya ideal untuk proyek-proyek konstruksi dan pertambangan.'),
('Motor Grader', 'Motor Grader adalah alat berat yang digunakan untuk meratakan dan membentuk permukaan jalan atau lahan. Dengan bilah peregangan di bagian tengahnya, motor grader dapat menghasilkan permukaan yang halus dan rata.'),
('Compactor', 'Compactor digunakan untuk memadatkan tanah, pasir, atau aspal. Alat ini membantu menciptakan permukaan yang kokoh dan stabil, khususnya dalam konstruksi jalan dan landasan.'),
('Wales', 'Wales, atau dikenal sebagai vibro roller, adalah alat berat yang menggunakan getaran untuk memadatkan tanah atau aspal. Wales efektif dalam menciptakan permukaan yang padat dan kuat.'),
('Scraper', 'Scraper digunakan untuk mengumpulkan, mengangkat, dan memindahkan material seperti tanah atau kerikil dalam jumlah besar. Scraper sangat efisien dalam pekerjaan penggalian dan pemindahan material.')
;
INSERT INTO equipments (name, category_id, description, price_per_day, price_per_week, price_per_month, price_per_year, available) VALUES
('Toyota 3 Ton Forklift', 1, 'Forklift diesel untuk gudang & industri', 900000, 6000000, 18000000, 185000000, true),
('Mitsubishi 5 Ton Forklift', 1, 'Cocok untuk angkat material berat outdoor', 1600000, 9500000, 28000000, 300000000, true),
('Tadano 25 Ton Truck Crane', 2, 'Crane mobile untuk proyek konstruksi', 6500000, 40000000, 120000000, 1200000000, true),
('Kato 50 Ton Crane', 2, 'Crane kapasitas menengah untuk proyek besar', 9000000, 58000000, 180000000, 1800000000, true),
('Sakai TW500W', 3, 'Roller tandem untuk finishing aspal', 1800000, 11000000, 32000000, 320000000, true),
('Caterpillar D6R', 4, 'Bulldozer pengurugan skala besar', 5500000, 35000000, 100000000, 1000000000, true),
('Komatsu D85EX Shovel', 5, 'Dozer dengan bucket untuk material loose', 6000000, 38000000, 115000000, 1150000000, true),
('Caterpillar 950H', 6, 'Loader serbaguna untuk material tambang', 4100000, 26000000, 78000000, 780000000, true),
('Caterpillar 140K', 7, 'Grader untuk jalan & land clearing', 4800000, 30000000, 90000000, 900000000, true),
('Bomag BW212D', 8, 'Compactor tanah vibratory 12 ton', 2900000, 18000000, 55000000, 550000000, true),
('Wacker Neuson BS50-2', 9, 'Tamping rammer untuk pemadatan spot', 200000, 1000000, 3000000, 30000000, true),
('Caterpillar 627K', 10, 'Scraper tandem loader', 8500000, 55000000, 160000000, 1600000000, true)
;




-- -- UPDATE RENTAL STATUS
-- CREATE OR REPLACE FUNCTION fn_update_rental_status()
-- RETURNS TRIGGER AS $$
-- BEGIN
--     IF OLD.status = 'BOOKED' AND NEW.start_date < NOW() THEN
--     NEW.status := 'ACTIVE';
--     ELSIF OLD.status = 'ACTIVE' AND NEW.end_date < NOW() THEN
--     NEW.status := 'COMPLETED';
--     END IF;
    
--     RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER trg_update_rental_status
-- BEFORE UPDATE ON rentals
-- FOR EACH ROW
-- EXECUTE FUNCTION fn_update_rental_status();

-- INSERT RENTAL HISTORIES
CREATE OR REPLACE FUNCTION fn_insert_rental_histories_status()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status = 'ACTIVE' THEN
        INSERT INTO rental_histories (rental_id, user_id, status, created_at) VALUES
        (NEW.id, NEW.user_id, 'ACTIVE', NOW());
    ELSIF NEW.status = 'COMPLETED' THEN
        INSERT INTO rental_histories (rental_id, user_id, status, created_at) VALUES
        (NEW.id, NEW.user_id, 'COMPLETED', NOW());
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_insert_rental_histories_status
AFTER UPDATE ON rentals
FOR EACH ROW
EXECUTE FUNCTION fn_insert_rental_histories_status();

CREATE OR REPLACE FUNCTION fn_update_equipment_availability()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status = 'COMPLETED' THEN
        UPDATE equipments SET available = true WHERE id = NEW.equipment_id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_update_equipment_availablity
AFTER UPDATE ON rentals
FOR EACH ROW
EXECUTE FUNCTION fn_update_equipment_availability();



-- DROP TRIGGER IF EXISTS trg_update_rental_status ON rentals;
DROP TRIGGER IF EXISTS trg_insert_rental_histories_status ON rentals;
DROP TRIGGER IF EXISTS trg_update_equipment_availablity ON rentals;
-- DROP FUNCTION IF EXISTS fn_update_rental_status();
DROP FUNCTION IF EXISTS fn_insert_rental_histories_status();
DROP FUNCTION IF EXISTS fn_update_equipment_availability();



DROP TABLE rental_histories;
DROP TABLE payments;
DROP TABLE rentals;
DROP TABLE equipments;
DROP TABLE categories;
DROP TABLE users;