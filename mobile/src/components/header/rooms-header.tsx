import React, { useState } from "react";
import { View, Text } from "native-base";
import { StyleSheet, Dimensions } from "react-native";
import { Colors } from "@/theme/variables";
import DropDownPicker from "react-native-dropdown-picker";

interface HeaderProps {

}

const screenWidth = Dimensions.get("screen").width;
const screenHeight = Dimensions.get("screen").height;

export const RoomsHeader: React.FC<HeaderProps> = () => {
  const [open, setOpen] = useState(false);
  const [selectedValue, setSelectedValue] = useState(null);
  const [items, setItems] = useState([
    { label: "Option 1", value: "option1" },
    { label: "Option 2", value: "option2" },
    { label: "Option 3", value: "option3" },
  ]);

  return (
    <View style={{ overflow: "hidden", paddingBottom: 1 }}>
      <View style={styles.dropdownContainer}>
        <DropDownPicker
          open={open}
          items={items}
          value={selectedValue}
          setOpen={setOpen}
          setValue={setSelectedValue}
          setItems={setItems}
          //   containerStyle={{ height: 40, width: 200 }}
          //   style={{ backgroundColor: "#fafafa" }}
          //   itemStyle={{
          //     justifyContent: "flex-start",
          //   }}
          //   dropDownStyle={{ backgroundColor: "#fafafa" }}
          //   onChangeItem={(item) => setSelectedValue(item.value)}
        />
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  dropdownContainer: {
    backgroundColor: Colors.WHITE,
    paddingHorizontal: 0.03 * screenWidth,
    paddingBottom: 0.02 * screenHeight,
    shadowColor: "#000",
    shadowOffset: { width: 1, height: 1 },
    shadowOpacity: 0.4,
    shadowRadius: 3,
    elevation: 5,
    // justifyContent: 'center'
  },
});
