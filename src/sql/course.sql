CREATE OR REPLACE FUNCTION update_field_on_zero()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.quantity = 0 AND OLD.quantity <> 0 THEN
        UPDATE racket
        SET avaliable = FALSE
        WHERE id = NEW.id;
    ELSEIF NEW.quantity <> 0 AND OLD.quantity = 0 THEN
        UPDATE racket
        SET avaliable = TRUE
        WHERE id = NEW.id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER trigger_zero_count
AFTER UPDATE ON racket
FOR EACH ROW
EXECUTE FUNCTION update_field_on_zero();
DROP TRIGGER trigger_zero_count ON racket;
