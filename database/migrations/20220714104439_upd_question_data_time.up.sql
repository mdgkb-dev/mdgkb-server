alter table questions
    alter column question_date type timestamp using question_date::timestamp;
