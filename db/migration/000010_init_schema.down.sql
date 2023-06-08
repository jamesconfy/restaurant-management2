DELETE FROM pg_type WHERE typname = 'payment_type';

DROP TABLE IF EXISTS payment_method;