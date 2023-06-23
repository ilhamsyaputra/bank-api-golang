-- install uuid module
create extension if not exists "uuid-ossp";

-- buat table
create table nasabah (
	no_nasabah serial not null,
	nama varchar(100) not null,
	nik varchar(55) not null,
	no_hp varchar(14) not null,
	tanggal_registrasi timestamp default now(),
	
	primary key(no_nasabah)
);

create table tipe_transaksi (
	kode_transaksi char(1),
	keterangan varchar(10),
	
	primary key(kode_transaksi)
);

create table rekening (
	no_rekening varchar(30) not null,
	no_nasabah int not null,
	saldo int not null,
	
	primary key(no_rekening),
	constraint no_nasabah
		foreign key(no_nasabah)
			references nasabah(no_nasabah)
);

create table transaksi (
	id uuid default uuid_generate_v4() not null,
	no_rekening varchar(30) not null,
	kode_transaksi char(1) not null,
	nominal int not null,
	waktu_transaksi timestamp default now(),
	
	primary key(id),
	
	constraint fk_no_rekening
		foreign key(no_rekening)
			references rekening(no_rekening),
			
	constraint fk_tipe
		foreign key(kode_transaksi)
			references tipe_transaksi(kode_transaksi)
);

insert into tipe_transaksi (kode_transaksi, keterangan) values 
	('D', 'Tarik'),
	('C', 'Tabung');

