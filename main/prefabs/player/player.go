components {
  id: "player_sprite"
  component: "/main/prefabs/player/player_sprite.sprite"
  position {
    x: -0.00547395
    y: 7.00189
  }
  scale {
    x: 1.362431
    y: 1.159354
    z: 1.0E-6
  }
}
components {
  id: "player"
  component: "/main/prefabs/player/scripts/player.script"
}
embedded_components {
  id: "player_collider"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 1.0\n"
  "friction: 0.1\n"
  "restitution: 0.0\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: -5.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 21.5\n"
  "  data: 45.5\n"
  "  data: 0.5\n"
  "}\n"
  "angular_damping: 1.0\n"
  ""
}
embedded_components {
  id: "flight_knife_factory"
  type: "factory"
  data: "prototype: \"/main/prefabs/knifes/flight_knife.go\"\n"
  ""
}
