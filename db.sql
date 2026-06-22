
-- tables --
CREATE TABLE IF NOT EXISTS Users(
	id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	owner_name VARCHAR(255) NOT NULL,
	balance BIGINT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	updated_at TIMESTAMPTZ NOT NULL,
	deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

	CONSTRAINT owner_name_not_short CHECK (LENGTH(TRIM(owner_name)) >= 2),
	CONSTRAINT balance_is_positive CHECK (balance >= 0)
); 

CREATE TABLE IF NOT EXISTS Transactions(
	id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	receiver_id BIGINT NOT NULL,
	sender_id BIGINT NOT NULL,
	status VARCHAR NOT NULL,
	amount BIGINT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	FOREIGN KEY (sender_id)	REFERENCES Users(id) ON DELETE SET NULL,
	FOREIGN KEY (receiver_id) REFERENCES Users(id) ON DELETE SET NULL,
	
	CONSTRAINT amount_is_positive CHECK (amount > 0),
	CONSTRAINT status_not_none CHECK (LENGTH(TRIM(status)) > 0)
);

CREATE OR REPLACE FUNCTION prevent_delete_with_balance()
RETURNS TRIGGER AS $$
BEGIN
	if OLD.balance <> 0 THEN
		RAISE EXCEPTION 'cannot delete account with balance above 0, (id=%, balance=%)',
			OLD.id, OLD.balance;

		END IF;
		RETURN OLD;
END;
$$ LANGUAGE plpgsql; 

CREATE TRIGGER trg_prevent_deletion_with_balance
	BEFORE DELETE ON Users
	FOR EACH ROW
	EXECUTE FUNCTION prevent_delete_with_balance();
