
CREATE OR REPLACE FUNCTION delete_records_on_false()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.avaliable = FALSE THEN
        DELETE FROM cart_racket
        WHERE racket_id = OLD.id;
    END IF;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER trigger_delete_on_false
AFTER UPDATE ON racket
FOR EACH ROW
EXECUTE FUNCTION delete_records_on_false();

DROP TRIGGER trigger_delete_on_false ON racket;