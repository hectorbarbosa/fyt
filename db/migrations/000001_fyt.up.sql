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

CREATE TABLE IF NOT EXISTS public.tasks (
    id BIGSERIAL PRIMARY KEY, 
    project_id BIGINT REFERENCES public.projects(id) NOT NULL, 
    title VARCHAR(50) NOT NULL,
    description VARCHAR NOT NULL,
    due_date TIMESTAMPTZ NOT NULL,
    doer INT REFERENCES public.users(id),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    done BOOLEAN
);