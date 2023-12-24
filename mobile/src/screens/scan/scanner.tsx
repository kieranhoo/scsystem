import React, { useState } from "react";
import {
  Dimensions,
  View,
  Text,
  StyleSheet,
  TouchableOpacity,
  TextInput,
  Modal,
  Image,
  TouchableWithoutFeedback,
  ActivityIndicator,
} from "react-native";
import { BarCodeScanner } from "expo-barcode-scanner";
import BarcodeMask from "react-native-barcode-mask";
import { Colors } from "@/theme/variables";
import { Header } from "@/components/header";
import { activity, room } from "@/services";
import { useFocusEffect } from "@react-navigation/native";

const screenWidth = Dimensions.get("screen").width;
const screenHeight = Dimensions.get("screen").height;

interface ResultPopupProps {
  isVisible: boolean;
  result: any | null;
  onClose: () => void;
  isLoading: boolean;
  openStatePopUp: any;
  setScanState: any;
  setMes: any;
  setLoading: any;
}
interface StatePopupProps {
  isVisible: boolean;
  result: boolean;
  onClose: () => void;
  mes: string;
  isLoading: boolean;
}
interface InputPopupProps {
  isVisible: boolean;
  onClose: () => void;
  check: any;
}

const ResultPopup: React.FC<ResultPopupProps> = ({
  isVisible,
  result,
  onClose,
  isLoading,
  setScanState,
  openStatePopUp,
  setMes,
  setLoading
}) => {
  const checkactivity = async () => {
    try {
      if (result.activity_type === "no access") {
        onClose()
        setMes("")
        setLoading(true)
        openStatePopUp(true)
        const checkresult = await activity.checkin(result.registration_id)
        setLoading(false)
        if (checkresult) {
          setMes(`Student ${result.last_name} ${result.first_name} check in successfully`)
          setScanState(true)
        }
        else {
          setMes(`Student ${result.last_name} ${result.first_name} check out fail`)
          setScanState(false)
        }
      }
      else if (result.activity_type === "in") {
        onClose()
        const checkresult = await activity.checkout(result.registration_id)
        if (checkresult) {
          setMes(`Student ${result.last_name} ${result.first_name} check out successfully`)
          setScanState(true)
          openStatePopUp(true)
        }
        else {
          setMes(`Student ${result.last_name} ${result.first_name} check out fail`)
          setScanState(false)
          openStatePopUp(true)
        }
      }
    }
    catch (err: any) {
      setMes(`Student ${result.last_name} ${result.first_name} check in/out fail`)
      setScanState(false)
      openStatePopUp(true)
    }
  }
  return (
    <Modal
      animationType="fade"
      transparent={true}
      visible={isVisible}
      onRequestClose={onClose}
    >
      <TouchableWithoutFeedback onPress={onClose}>
        <View style={styles.modalContainer}>
          <View style={styles.modalContent}>
            {
              isLoading
                ?
                <View>
                  <ActivityIndicator size="large" color={Colors.PRIMARY} />
                  <Text style={{ marginTop: 10, textAlign: "center", fontFamily: "Poppins_400Regular", }}>Please wait for checking...</Text>
                </View>
                :
                <View>
                  <View style={styles.studentInfor}>
                    <Text style={styles.studentName}>
                      {result.last_name} {result.first_name}
                    </Text>
                    <Text style={styles.studentSpecialInfor}>
                      ID: {result.student_id}
                    </Text>
                    <Text style={styles.studentSpecialInfor}>
                      Email: {result.email}
                    </Text>
                  </View>
                  <View>
                    <View style={styles.studentInfor}>
                      <Text style={styles.studentSubInfor}>
                        Registeration time: {result.registration_time}
                      </Text>
                    </View>
                    <View>
                      <Text style={styles.studentSubInfor}>
                        Room: {result.room_name}
                      </Text>
                    </View>
                    <View>
                      <Text style={styles.studentSubInfor}>
                        Supervisor: {result.supervisor}
                      </Text>
                    </View>
                  </View>
                  <View style={result.activity_type !== "out" ? styles.buttonList : styles.buttonSingle}>
                    {
                      result.activity_type !== "out" &&
                      // <Button color={Colors.PRIMARY} title={`${result.activity_type === "no access" ? "Check in" : "Check out"}`} onPress={() => checkactivity()} />
                      <TouchableOpacity
                        onPress={() => checkactivity()}
                        style={styles.button}
                      >
                        <Text style={styles.buttontext}>
                          {result.activity_type === "no access" ? "Check in" : "Check out"}
                        </Text>
                      </TouchableOpacity>
                    }
                    {/* <Button color={Colors.PRIMARY} title="CLOSE" onPress={onClose} /> */}
                    <TouchableOpacity
                      onPress={onClose}
                      style={styles.cancelButton}
                    >
                      <Text style={styles.cancelButtonText}>
                        CLOSE
                      </Text>
                    </TouchableOpacity>
                  </View>
                </View>
            }
          </View>
        </View>
      </TouchableWithoutFeedback>
    </Modal>
  );
};

const StatePopup: React.FC<StatePopupProps> = ({
  isVisible,
  result,
  onClose,
  mes,
  isLoading
}) => {
  return (
    <Modal
      animationType="fade"
      transparent={true}
      visible={isVisible}
      onRequestClose={onClose}
    >
      <TouchableWithoutFeedback onPress={onClose}>
        <View style={styles.modalContainer}>
          <View style={styles.modalContent}>
            {isLoading ? (
              <View>
                <ActivityIndicator size="large" color={Colors.PRIMARY} />
                <Text style={{ marginTop: 10, fontFamily: "Poppins_400Regular" }}>
                  {mes}
                </Text>
              </View>
            ) : (
              <View style={styles.imageContainer}>
                <Image
                  style={styles.image}
                  source={
                    result
                      ? require("../../../assets/scan/success.png")
                      : require("../../../assets/scan/fail.png")
                  }
                />
                <Text
                  style={
                    result ? styles.stateSuccess : styles.stateError
                  }
                >
                  {result ? "Success" : "Error"}
                </Text>
                <Text style={styles.statusText}>{mes}</Text>
              </View>
            )}
          </View>
        </View>
      </TouchableWithoutFeedback>
    </Modal>
  );
};

const InputPopup: React.FC<InputPopupProps> = ({ isVisible, onClose, check }) => {
  const [value, setValue] = useState("")
  const handleClick = () => {
    onClose();
    check({ data: value })
  }
  return (
    <Modal
      animationType="fade"
      transparent={true}
      visible={isVisible}
      onRequestClose={onClose}
    >
      <TouchableWithoutFeedback onPress={onClose}>
        <View style={styles.modalContainer}>
          <View style={styles.modalContent}>
            <Text style={styles.label}>Student ID:</Text>
            <TextInput
              placeholder="Type Student ID to check in/out"
              value={value}
              onChangeText={setValue}
              style={styles.input}
              keyboardType="numeric"
            />
            {/* <Button color={Colors.PRIMARY} title="OK" onPress={onClose} /> */}
            <TouchableOpacity
              onPress={handleClick}
              style={styles.button}
            >
              <Text style={styles.buttontext}>OK</Text>
            </TouchableOpacity>
          </View>
        </View>
      </TouchableWithoutFeedback>
    </Modal>
  );
};

export const Scanner = () => {
  const [hasPermission, setHasPermission] = useState<boolean | null>(null);
  const [scanned, setScanned] = useState(false);
  const [result, setResult] = useState<string | null>(null);
  const [showPopup, setShowPopup] = useState(false);
  const [showInputPopup, setShowInputPopup] = useState(false);
  const [showState, setShowState] = useState(false);
  const [cameraBackground] = useState("transparent");
  const [requestingPermission, setRequestingPermission] = useState(true);
  const [isLoading, setIsLoading] = useState(false)
  const [scanstate, setScanState] = useState(false)
  const [statemes, setStateMes] = useState("")

  const handleBarCodeScanned = async ({ data }: { data: string }) => {
    try {
      setIsLoading(true)
      if (isNaN(parseInt(data))) {
        setScanState(false)
        setStateMes("Scan student ID fail")
        setShowState(true)
        setIsLoading(false)
      }
      else {
        setShowPopup(true);
        const checkresult = await room.check(data, "1");
        setResult(checkresult.data);
        setIsLoading(false)
      }
    }
    catch (err: any) {
      setScanned(true);
      setScanState(false)
      setStateMes("Scan student ID fail")
      setShowState(true)
      setIsLoading(false)
    }
  };

  const handleClosePopup = () => {
    setScanned(false);
    setResult(null);
    setShowPopup(false);
  };

  const handleCloseInputPopup = () => {
    setScanned(false);
    setShowInputPopup(false);
  };

  const handleCloseState = () => {
    setScanned(false);
    setShowState(false);
  };

  const handleShowInputPopUp = () => {
    setScanned(true);
    setShowInputPopup(true)
  }

  // useEffect(() => {
  //   (async () => {
  //     setRequestingPermission(true); // Bắt đầu xin quyền
  //     const { status } = await BarCodeScanner.requestPermissionsAsync();
  //     setHasPermission(status === "granted");
  //     setRequestingPermission(false); // Kết thúc xin quyền
  //   })();
  // }, []);

  useFocusEffect(
    React.useCallback(() => {
      (async () => {
        setRequestingPermission(true); // Bắt đầu xin quyền
        const { status } = await BarCodeScanner.requestPermissionsAsync();
        setHasPermission(status === "granted");
        setRequestingPermission(false); // Kết thúc xin quyền
      }
      )();
      return () => {
        setRequestingPermission(false)
        setHasPermission(null)
        setScanned(false);
        setResult(null);
        setShowPopup(false);
      };
    }, [])
  );

  return (
    <View style={styles.container}>
      <View>
        <Header title="Scan" />
      </View>
      {requestingPermission && (
        <View>
        </View>
        // <Modal transparent={true} visible={true}>
        //   <View style={styles.modalContainer}>
        //     <View style={styles.modalContent}>
        //       <ActivityIndicator size="large" color={Colors.PRIMARY} />
        //       <Text style={{ marginTop: 10, fontFamily: "Poppins_400Regular" }}>Requesting camera permission...</Text>
        //     </View>
        //   </View>
        // </Modal>
      )
      }
      {
        hasPermission !== null ?
          <>
            {
              scanned ? <BarCodeScanner
                onBarCodeScanned={undefined}
                style={styles.camera}
              >
                <View style={styles.cameraMask}>
                  <BarcodeMask
                    width={300}
                    height={300}
                    edgeWidth={20}
                    edgeHeight={20}
                    edgeColor={Colors.PRIMARY}
                    edgeRadius={10}
                    edgeBorderWidth={5}
                    animatedLineColor={Colors.PRIMARY}
                    animatedLineHeight={1.5}
                    animatedLineWidth={200}
                    lineAnimationDuration={2000}
                    useNativeDriver
                    backgroundColor={cameraBackground}
                  />
                </View>
              </BarCodeScanner>
                :
                <BarCodeScanner
                  onBarCodeScanned={handleBarCodeScanned}
                  style={styles.camera}
                >
                  <View style={styles.cameraMask}>
                    <BarcodeMask
                      width={300}
                      height={300}
                      edgeWidth={20}
                      edgeHeight={20}
                      edgeColor={Colors.PRIMARY}
                      edgeRadius={10}
                      edgeBorderWidth={5}
                      animatedLineColor={Colors.PRIMARY}
                      animatedLineHeight={1.5}
                      animatedLineWidth={200}
                      lineAnimationDuration={2000}
                      useNativeDriver
                      backgroundColor={cameraBackground}
                    />
                  </View>
                </BarCodeScanner>
            }
            <View style={styles.inputcontainer}>
              <TouchableOpacity
                onPress={handleShowInputPopUp}
                style={styles.button}
              >
                <Text style={styles.buttontext}>Add Record</Text>
              </TouchableOpacity>
            </View>
            {showInputPopup && (
              <InputPopup isVisible={showInputPopup} onClose={handleCloseInputPopup} check={handleBarCodeScanned} />
            )}
            {showPopup && (
              <ResultPopup isVisible={showPopup} result={result} isLoading={isLoading} onClose={handleClosePopup} setScanState={setScanState} openStatePopUp={setShowState} setMes={setStateMes} setLoading={setIsLoading} />
            )}
            {showState && (
              <StatePopup isVisible={showState} result={scanstate} onClose={handleCloseState} mes={statemes} isLoading={isLoading} />
            )}
          </> : <></>
      }
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    height: screenHeight,
    backgroundColor: Colors.WHITE,
  },
  modalContainer: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: "rgba(0,0,0,0.5)",
  },
  camera: {
    height: "70%",
    alignSelf: "center",
    width: screenWidth,
  },
  cameraMask: {
    flex: 1,
  },
  modalContent: {
    backgroundColor: "white",
    padding: 16,
    borderRadius: 10,
    elevation: 5,
    width: 0.8 * screenWidth,
  },
  button: {
    backgroundColor: Colors.PRIMARY,
    paddingHorizontal: 10,
    paddingVertical: 4,
    borderRadius: 10,
    justifyContent: "center",
  },
  buttontext: {
    color: Colors.WHITE,
    fontFamily: "Poppins_400Regular",
    textAlign: "center",
    textAlignVertical: "center"
  },
  input: {
    padding: 4,
    marginBottom: 12,
    borderWidth: 1,
    borderColor: "gray",
    borderRadius: 10,
  },
  inputcontainer: {
    alignItems: "center",
    justifyContent: "center",
    padding: 5,
    fontFamily: "Poppins_400Regular",
  },
  label: {
    fontFamily: "Poppins_600SemiBold",
  },
  studentInfor: {
    width: '100%',
    textAlign: "center"
  },
  studentName: {
    fontFamily: "Poppins_600SemiBold",
    fontSize: 25,
    textAlign: "center"
  },
  buttonList: {
    flexDirection: "row",
    marginTop: 8,
    justifyContent: "space-between",
    alignContent: "center",
  },
  buttonSingle: {
    flexDirection: "row",
    marginTop: 8,
    justifyContent: "center",
    alignContent: "center",
  },
  studentSpecialInfor: {
    fontFamily: "Poppins_600SemiBold",
    fontSize: 18,
    color: Colors.PRIMARY,
    textAlign: "center"
  },
  studentSubInfor: {
    marginTop: 4,
    color: Colors.GREY,
    fontFamily: "Poppins_400Regular",
  },
  cancelButton: {
    borderColor: Colors.PRIMARY,
    borderRadius: 10,
    borderWidth: 2,
    paddingHorizontal: 10,
    paddingVertical: 4,
    justifyContent: "center",
  },
  cancelButtonText: {
    color: Colors.PRIMARY,
    fontFamily: "Poppins_400Regular",
    textAlign: "center",
    textAlignVertical: "center"
  },
  imageContainer: {
    alignItems: 'center',
    justifyContent: 'center',
  },

  image: {
    width: 70,
    height: 70,
  },

  statusText: {
    marginTop: 10,
    textAlign: 'center',
    fontFamily: "Poppins_400Regular",
    color: Colors.GREY
  },
  stateSuccess: {
    fontSize: 20,
    textAlign: "center",
    color: Colors.SUCCESS,
    marginBottom: 4,
    fontFamily: "Poppins_400Regular",
  },
  stateError: {
    fontSize: 20,
    textAlign: "center",
    color: Colors.ERROR,
    marginBottom: 4,
    fontFamily: "Poppins_400Regular",
  }
});
