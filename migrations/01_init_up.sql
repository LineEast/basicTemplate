create table users (
	id uuid default gen_random_uuid(),
	name varchar(32) not null,
	phoneNumber varchar(16) not null
);
