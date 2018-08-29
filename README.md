# Gitlab CI Validate

Validate a .gitlab-ci.yml file against Gitlab

## Running the program

**Valid .gitlab-ci.yml**

```
$ gitlab-ci-validate --filepath ./.gitlab-ci.yml
Valid gitlab-ci.yml file
```

**Invalid .gitlab-ci.yml**

```
$ gitlab-ci-validate --filepath ./.gitlab-ci.yml
Invalid contents:
Error: Invalid configuration format
```
