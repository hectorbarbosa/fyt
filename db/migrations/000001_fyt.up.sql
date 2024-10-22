CREATE TABLE IF NOT EXISTS public.projects (
    id BIGSERIAL PRIMARY KEY, 
    project_type INT NOT NULL,
    title VARCHAR(50) NOT NULL,
    description VARCHAR NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    social_url VARCHAR [],
    source_url VARCHAR,
    closed BOOLEAN DEFAULT false,
    closed_at TIMESTAMPTZ
)