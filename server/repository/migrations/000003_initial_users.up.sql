INSERT INTO "user"
    (username, email, given, surname, phone, is_active, is_admin, is_login_disabled, is_verified, is_mfa_enabled)
VALUES
    ('deeprave', 'davidn@uniquode.io', 'David', 'Nugent', '+61404867638', 't', 't', 'f', 't', 't' ),
    ('devtest', 'uniquode.dev@gmail.com', 'Developer', 'Account', '+61404867638', 't', 'f', 'f', 't', 'f'),
    ('testdev', 'invalid@gmail.com', 'Invalid', 'Account', NULL, 't', 'f', 't', 't', 'f'),
    ('invalid', NULL, 'Another', 'Invalid', NULL, 't', 'f', 't', 't', 'f'),
    ('unverified', NULL, 'Unverified', 'User', NULL, 't', 'f', 't', 't', 'f')
;
