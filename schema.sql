--
-- PostgreSQL database dump
--

-- Dumped from database version 11.3
-- Dumped by pg_dump version 11.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: ltree; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS ltree WITH SCHEMA public;


--
-- Name: EXTENSION ltree; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION ltree IS 'data type for hierarchical tree-like structures';


--
-- Name: update_categories_parent_path(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_categories_parent_path() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    DECLARE
        path ltree;
    BEGIN
        IF NEW.parent IS NULL THEN
            NEW.parent_path = 'root'::ltree;
        ELSEIF TG_OP = 'INSERT' OR OLD.parent IS NULL OR OLD.parent != NEW.parent THEN
            SELECT parent_path || id::text FROM section WHERE id = NEW.parent INTO path;
            IF path IS NULL THEN
                RAISE EXCEPTION 'Invalid parent_id %', NEW.parent_id;
            END IF;
            NEW.parent_path = path;
        END IF;
        RETURN NEW;
    END;
$$;


ALTER FUNCTION public.update_categories_parent_path() OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: answer; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.answer (
    id integer NOT NULL,
    "user" integer NOT NULL,
    survey integer NOT NULL,
    score smallint,
    body character varying
);


ALTER TABLE public.answer OWNER TO postgres;

--
-- Name: answer_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.answer_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.answer_id_seq OWNER TO postgres;

--
-- Name: answer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.answer_id_seq OWNED BY public.answer.id;


--
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying,
    parent integer,
    parent_path public.ltree
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: TABLE categories; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.categories IS 'category hierarchy path (ref: https://coderwall.com/p/whf3-a/hierarchical-data-in-postgres)';


--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.categories_id_seq OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: notes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.notes (
    id integer NOT NULL,
    text text,
    saved timestamp with time zone NOT NULL
);


ALTER TABLE public.notes OWNER TO postgres;

--
-- Name: notes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.notes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.notes_id_seq OWNER TO postgres;

--
-- Name: notes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.notes_id_seq OWNED BY public.notes.id;


--
-- Name: orgs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orgs (
    id integer NOT NULL,
    name character varying
);


ALTER TABLE public.orgs OWNER TO postgres;

--
-- Name: orgs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.orgs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.orgs_id_seq OWNER TO postgres;

--
-- Name: orgs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.orgs_id_seq OWNED BY public.orgs.id;


--
-- Name: question; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.question (
    id integer NOT NULL,
    open boolean,
    body text NOT NULL
);


ALTER TABLE public.question OWNER TO postgres;

--
-- Name: question_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.question_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.question_id_seq OWNER TO postgres;

--
-- Name: question_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.question_id_seq OWNED BY public.question.id;


--
-- Name: survey; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.survey (
    id integer NOT NULL,
    start date,
    "end" date,
    score_question integer,
    open_question integer,
    note integer
);


ALTER TABLE public.survey OWNER TO postgres;

--
-- Name: survey_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.survey_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.survey_id_seq OWNER TO postgres;

--
-- Name: survey_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.survey_id_seq OWNED BY public.survey.id;


--
-- Name: tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tags (
    id integer NOT NULL,
    name character varying,
    attribute character varying,
    number integer,
    "timestamp" timestamp without time zone
);


ALTER TABLE public.tags OWNER TO postgres;

--
-- Name: tags_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tags_id_seq OWNER TO postgres;

--
-- Name: tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tags_id_seq OWNED BY public.tags.id;


--
-- Name: user_category_idx; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_category_idx (
    id integer NOT NULL,
    "user" integer,
    category integer
);


ALTER TABLE public.user_category_idx OWNER TO postgres;

--
-- Name: user_category_idx_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_category_idx_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_category_idx_id_seq OWNER TO postgres;

--
-- Name: user_category_idx_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_category_idx_id_seq OWNED BY public.user_category_idx.id;


--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq OWNER TO postgres;

--
-- Name: user_org_idx; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_org_idx (
    id integer NOT NULL,
    "user" integer,
    org integer
);


ALTER TABLE public.user_org_idx OWNER TO postgres;

--
-- Name: user_org_idx_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_org_idx_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_org_idx_id_seq OWNER TO postgres;

--
-- Name: user_org_idx_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_org_idx_id_seq OWNED BY public.user_org_idx.id;


--
-- Name: user_tag_idx; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_tag_idx (
    id integer NOT NULL,
    "user" integer,
    tag integer
);


ALTER TABLE public.user_tag_idx OWNER TO postgres;

--
-- Name: user_tag_idx_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_tag_idx_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_tag_idx_id_seq OWNER TO postgres;

--
-- Name: user_tag_idx_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_tag_idx_id_seq OWNED BY public.user_tag_idx.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying,
    first_name character varying,
    last_name character varying,
    nick_name character varying
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: answer id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answer ALTER COLUMN id SET DEFAULT nextval('public.answer_id_seq'::regclass);


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: notes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notes ALTER COLUMN id SET DEFAULT nextval('public.notes_id_seq'::regclass);


--
-- Name: orgs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orgs ALTER COLUMN id SET DEFAULT nextval('public.orgs_id_seq'::regclass);


--
-- Name: question id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.question ALTER COLUMN id SET DEFAULT nextval('public.question_id_seq'::regclass);


--
-- Name: survey id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.survey ALTER COLUMN id SET DEFAULT nextval('public.survey_id_seq'::regclass);


--
-- Name: tags id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags ALTER COLUMN id SET DEFAULT nextval('public.tags_id_seq'::regclass);


--
-- Name: user_category_idx id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_category_idx ALTER COLUMN id SET DEFAULT nextval('public.user_category_idx_id_seq'::regclass);


--
-- Name: user_org_idx id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_org_idx ALTER COLUMN id SET DEFAULT nextval('public.user_org_idx_id_seq'::regclass);


--
-- Name: user_tag_idx id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_tag_idx ALTER COLUMN id SET DEFAULT nextval('public.user_tag_idx_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: answer answer_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answer
    ADD CONSTRAINT answer_pk PRIMARY KEY (id);


--
-- Name: categories categories_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pk PRIMARY KEY (id);


--
-- Name: notes notes_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notes
    ADD CONSTRAINT notes_pk PRIMARY KEY (id);


--
-- Name: orgs orgs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orgs
    ADD CONSTRAINT orgs_pk PRIMARY KEY (id);


--
-- Name: question question_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.question
    ADD CONSTRAINT question_pk PRIMARY KEY (id);


--
-- Name: survey survey_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.survey
    ADD CONSTRAINT survey_pk PRIMARY KEY (id);


--
-- Name: tags tags_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pk PRIMARY KEY (id);


--
-- Name: user_category_idx user_category_idx_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_category_idx
    ADD CONSTRAINT user_category_idx_pk PRIMARY KEY (id);


--
-- Name: user_org_idx user_org_idx_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_org_idx
    ADD CONSTRAINT user_org_idx_pk PRIMARY KEY (id);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- Name: answer_survey_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX answer_survey_idx ON public.answer USING btree (survey);


--
-- Name: answer_user_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX answer_user_idx ON public.answer USING btree ("user");


--
-- Name: categories_name_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX categories_name_idx ON public.categories USING btree (name);


--
-- Name: categories_parent_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX categories_parent_idx ON public.categories USING btree (parent);


--
-- Name: tags_name_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX tags_name_idx ON public.tags USING btree (name);


--
-- Name: user_org_idx_user_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX user_org_idx_user_idx ON public.user_org_idx USING btree ("user");


--
-- Name: user_tag_idx_tag_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX user_tag_idx_tag_idx ON public.user_tag_idx USING btree (tag);


--
-- Name: user_tag_idx_user_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX user_tag_idx_user_idx ON public.user_tag_idx USING btree ("user");


--
-- Name: categories parent_path_tgr; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER parent_path_tgr BEFORE INSERT OR UPDATE ON public.categories FOR EACH ROW EXECUTE PROCEDURE public.update_categories_parent_path();


--
-- Name: answer answer_survey_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answer
    ADD CONSTRAINT answer_survey_fk FOREIGN KEY (survey) REFERENCES public.survey(id) ON DELETE RESTRICT;


--
-- Name: answer answer_user_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.answer
    ADD CONSTRAINT answer_user_fk FOREIGN KEY ("user") REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: categories categories_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_fk FOREIGN KEY (parent) REFERENCES public.categories(id) ON DELETE CASCADE;


--
-- Name: survey survey_note_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.survey
    ADD CONSTRAINT survey_note_fk FOREIGN KEY (note) REFERENCES public.notes(id) ON DELETE RESTRICT;


--
-- Name: survey survey_open_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.survey
    ADD CONSTRAINT survey_open_fk FOREIGN KEY (open_question) REFERENCES public.question(id) ON DELETE RESTRICT;


--
-- Name: survey survey_score_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.survey
    ADD CONSTRAINT survey_score_fk FOREIGN KEY (score_question) REFERENCES public.question(id) ON DELETE RESTRICT;


--
-- Name: user_category_idx user_category_category_idx_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_category_idx
    ADD CONSTRAINT user_category_category_idx_fk FOREIGN KEY (category) REFERENCES public.categories(id) ON DELETE CASCADE;


--
-- Name: user_category_idx user_category_user_idx_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_category_idx
    ADD CONSTRAINT user_category_user_idx_fk FOREIGN KEY ("user") REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: user_org_idx user_org_org_idx_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_org_idx
    ADD CONSTRAINT user_org_org_idx_fk FOREIGN KEY (org) REFERENCES public.orgs(id) ON DELETE CASCADE;


--
-- Name: user_org_idx user_org_user_idx_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_org_idx
    ADD CONSTRAINT user_org_user_idx_fk FOREIGN KEY ("user") REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: user_tag_idx user_tag_idx_tag_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_tag_idx
    ADD CONSTRAINT user_tag_idx_tag_fk FOREIGN KEY (tag) REFERENCES public.tags(id) ON DELETE CASCADE;


--
-- Name: user_tag_idx user_tag_idx_user_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_tag_idx
    ADD CONSTRAINT user_tag_idx_user_fk FOREIGN KEY ("user") REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

