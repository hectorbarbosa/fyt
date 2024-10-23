CREATE TABLE IF NOT EXISTS public.users (
    id SERIAL PRIMARY KEY, 
    email VARCHAR NOT NULL,
    user_name VARCHAR(50) NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS public.projects (
    id BIGSERIAL PRIMARY KEY, 
    owner INT REFERENCES public.users(id) NOT NULL,
    project_type INT NOT NULL,
    title VARCHAR(50) NOT NULL,
    description VARCHAR NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    social_url VARCHAR [],
    source_url VARCHAR,
    closed BOOLEAN DEFAULT false,
    closed_at TIMESTAMPTZ
);

