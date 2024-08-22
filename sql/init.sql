--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: faces; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.faces (
    width integer NOT NULL,
    height integer NOT NULL,
    x integer NOT NULL,
    y integer NOT NULL,
    gender character varying(16) NOT NULL,
    age integer NOT NULL,
    image_id integer
);


ALTER TABLE public.faces OWNER TO admin;

--
-- Name: images; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.images (
    id integer NOT NULL,
    task_id character varying(128),
    title character varying(256) NOT NULL
);


ALTER TABLE public.images OWNER TO admin;

--
-- Name: images_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.images_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.images_id_seq OWNER TO admin;

--
-- Name: images_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.images_id_seq OWNED BY public.images.id;


--
-- Name: tasks; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.tasks (
    uuid character varying(128) NOT NULL,
    progress character varying(128) NOT NULL
);


ALTER TABLE public.tasks OWNER TO admin;

--
-- Name: images id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.images ALTER COLUMN id SET DEFAULT nextval('public.images_id_seq'::regclass);


--
-- Data for Name: faces; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.faces (width, height, x, y, gender, age, image_id) FROM stdin;
\.


--
-- Data for Name: images; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.images (id, task_id, title) FROM stdin;
\.


--
-- Data for Name: tasks; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.tasks (uuid, progress) FROM stdin;
\.


--
-- Name: images_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.images_id_seq', 49, true);


--
-- Name: images images_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_pkey PRIMARY KEY (id);


--
-- Name: tasks tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (uuid);


--
-- Name: faces faces_image_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.faces
    ADD CONSTRAINT faces_image_id_fkey FOREIGN KEY (image_id) REFERENCES public.images(id) ON DELETE CASCADE;


--
-- Name: images images_task_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_task_id_fkey FOREIGN KEY (task_id) REFERENCES public.tasks(uuid) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

