CREATE TABLE IF NOT EXISTS "user" (
    id bigserial PRIMARY KEY,
    username varchar (63) UNIQUE NOT NULL,
    email varchar (127) UNIQUE,
    given varchar(63),
    surname varchar(63),
    phone varchar(50),
    is_active bool DEFAULT 'f' NOT NULL,
    is_admin bool DEFAULT 'f' NOT NULL,
    is_login_disabled bool DEFAULT 'f' NOT NULL,
    is_verified bool DEFAULT 'f' NOT NULL,
    is_mfa_enabled bool DEFAULT 'f' NOT NULL,
    dt_lastlogin timestamp,
    dt_created timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    dt_updated timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    dt_deleted timestamp DEFAULT NULL
);

