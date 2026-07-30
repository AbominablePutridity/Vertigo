components {
  id: "flight_knife"
  component: "/main/prefabs/knifes/flight_knife.sprite"
  position {
    x: -0.24647707
    y: 2.7859554
  }
  scale {
    x: 1.400417
    y: 1.358767
  }
}
components {
  id: "flight_knife1"
  component: "/main/prefabs/flight_weapon.script"
}
components {
  id: "physic_throwing"
  component: "/main/prefabs/physic_throwing.script"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 0.5\n"
  "friction: 0.05\n"
  "restitution: 0.7\n"
  "group: \"knife\"\n"
  "mask: \"default\"\n"
  "mask: \"enemy\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: 2.5\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 13.625498\n"
  "  data: 2.478612\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
