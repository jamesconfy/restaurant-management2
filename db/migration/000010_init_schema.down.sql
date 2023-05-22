DELETE FROM pg_type WHERE typname = 'payment_status';

DROP TABLE IF EXISTS payments;