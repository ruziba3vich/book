DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM   information_schema.tables 
        WHERE  table_schema = 'public'
        AND    table_name = 'authors'
    ) THEN
        CREATE TABLE authors (
            id UUID,
            fullname VARCHAR(100),
            PRIMARY KEY (id)
        );

    END IF;
END$$;
