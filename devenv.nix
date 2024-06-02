{ pkgs, lib, config, inputs, ... }:

{
  env.GREET = "devenv";

  languages.javascript.enable = true;
  languages.typescript.enable = true;

  packages = [
    pkgs.git
    pkgs.graphviz
    pkgs.devenv

    pkgs.age

    pkgs.nodePackages.npm
    pkgs.nodePackages.typescript-language-server
    pkgs.nodePackages.svelte-language-server
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
