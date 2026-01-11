DROP TABLE IF EXISTS user_data;
CREATE TABLE user_data (
    uuid TEXT PRIMARY KEY,
    u_seq INTEGER,
    u_id TEXT,
    u_name TEXT,
    u_nick TEXT,
    u_cp INTEGER,
    u_candy REAL,
    u_like REAL,
    u_fans INTEGER,
    u_fans_grade INTEGER,
    u_selected_costume_id INTEGER,
    u_selected_music_id INTEGER,
    u_retain_continuous_tap INTEGER,
    u_join_type INTEGER,
    u_last_login TEXT,
    u_last_communication TEXT,
    u_save_date TEXT,
    u_create_time TEXT,
    u_tutorial_step INTEGER,
    u_review_popup TEXT,
    device_uuid TEXT,
    u_samseck_step INTEGER,
    u_free_cp INTEGER,
    u_charge_cp INTEGER
);
DROP TABLE IF EXISTS user_area_info;
CREATE TABLE user_area_info (
    uuid TEXT NOT NULL,
    u_area_num INTEGER NOT NULL,
    d_like REAL,
    i_user_fan_count INTEGER,
    i_user_fan_grade INTEGER,
    i_selected_costume_id INTEGER,
    i_selected_music_id INTEGER,
    i_selected_guitar_id INTEGER,
    d_candy REAL,
    s_tutorial_list TEXT,
    s_gp1 TEXT,
    s_gp2 TEXT,
    PRIMARY KEY (uuid, u_area_num),
    FOREIGN KEY (uuid) REFERENCES user_data(uuid) ON DELETE CASCADE
);
DROP TABLE IF EXISTS user_achievement;
CREATE TABLE user_achievement (
    uuid TEXT NOT NULL,
    i_id: INTEGER NOT NULL,
    i_level: INTEGER,
    d_quantity: INTEGER,
    s_quantity: TEXT,
    PRIMARY KEY (uuid, i_id),
    FOREIGN KEY (uuid) REFERENCES user_data(uuid) ON DELETE CASCADE
)