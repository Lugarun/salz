{ pkgs, lib, config, inputs, ... }:

{
  env.GREET = "devenv";

  packages = [
    pkgs.git
    pkgs.graphviz
    pkgs.devenv

    pkgs.age
  ];

  services.postgres = {
    enable = true;
    initialDatabases = [{ name = "salz"; schema = ./db/init.sql; }];
    initialScript = ''
      CREATE ROLE salz WITH LOGIN PASSWORD 'superdupersecret';
      ALTER ROLE salz WITH SUPERUSER; -- TODO CHANGE THIS IN PROD LOL
    '';
    listen_addresses = "127.0.0.1";
  };

  services.minio = {
    enable = true;
    consoleAddress = "127.0.0.1:9111";
    listenAddress = "127.0.0.1:9110";
    buckets = ["salz"];
  };
  services.caddy.enable = true;

  languages.go.enable = true;
}
