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
    start_date date default current_date not null,
    end_date date not null,
    price decimal(12, 2) not null,
    status varchar(20) not null,
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

