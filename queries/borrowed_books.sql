DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM   information_schema.tables 
        WHERE  table_schema = 'public'
        AND    table_name = 'borrowed_books'
    ) THEN
        CREATE TABLE borrowed_books (
            id SERIAL PRIMARY KEY,
            book_id UUID,
            user_id UUID,
            FOREIGN KEY (book_id) REFERENCES books(id),
            FOREIGN KEY (user_id) REFERENCES users(id)
        );
    END IF;
END$$;
