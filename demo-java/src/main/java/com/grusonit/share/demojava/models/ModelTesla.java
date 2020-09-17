package com.grusonit.share.demojava.models;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;

@Entity
public class ModelTesla {
    @Id
    @GeneratedValue
    public long id;
    public String name;
    public String type;
    public String description;
    public String photoLink;
}
