DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM   information_schema.tables 
        WHERE  table_schema = 'public'
        AND    table_name = 'books'
    ) THEN
        CREATE TABLE books (
            id UUID,
            name VARCHAR(100),
            author UUID,
            PRIMARY KEY (id),
            FOREIGN KEY (author) REFERENCES authors(id)
        );
    END IF;
END$$;
