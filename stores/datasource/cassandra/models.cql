USE stores;
CREATE TABLE IF NOT EXISTS store_login (
    admin_user_email text PRIMARY KEY,
    password text
);

CREATE TABLE IF NOT EXISTS store_information (
    id uuid PRIMARY KEY,
    admin_user_name text,
    admin_user_last_name text,
    admin_user_email text,
    password text,
    created_at text,
    updated_at text,
    permissions frozen<Permissions>,
    role frozen<StoreAddress>
);

CREATE TABLE IF NOT EXISTS store_address (
    id uuid PRIMARY KEY,
    admin_user boolean,
    role_inventory_read boolean,
    role_super_admin boolean
);

CREATE TABLE IF NOT EXISTS permissions (
    id uuid PRIMARY KEY,
    admin_user_id text,
    read_stores boolean,
    read_write_stores boolean,
    read_inventory boolean,
    read_write_inventory boolean,
    read_admin_user boolean,
    read_write_admin_user boolean,
    read_config boolean,
    read_write_config boolean
);