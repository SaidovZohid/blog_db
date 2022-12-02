CREATE TABLE IF NOT EXISTS "comments"(
    "id" SERIAL PRIMARY KEY,
    "post_id" int not null REFERENCES posts(id),
    "user_id" int not null REFERENCES users(id),
    "description" text not null,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS "likes"(
    "id" SERIAL PRIMARY KEY,
    "post_id" int not null REFERENCES users(id),
    "user_id" int not null REFERENCES users(id),
    "status" boolean not null,
    UNIQUE(post_id, user_id)
);