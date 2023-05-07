CREATE TABLE IF NOT EXISTS event (
  id UUID PRIMARY KEY,
  title TEXT NOT NULL,
  is_deleted BOOLEAN DEFAULT FALSE);
CREATE INDEX IF NOT EXISTS idx__event__title ON event(title);
-- 
CREATE TABLE IF NOT EXISTS "user" (id UUID PRIMARY KEY);
--
CREATE TABLE IF NOT EXISTS event_user (
  id SERIAL PRIMARY KEY,
  event_id UUID REFERENCES event(id) ON DELETE CASCADE NOT NULL,
  user_id UUID REFERENCES "user"(id) ON DELETE CASCADE NOT NULL);
CREATE INDEX IF NOT EXISTS idx__event_user__event_id ON event_user(event_id);
CREATE INDEX IF NOT EXISTS idx__event_user__user_id ON event_user(user_id);
